package gittp

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
)

type router struct {
	check func(RequestInfo) (bool, int)
}

// IsGitRequest is a helper to ensure this is a git request
func IsGitRequest(path string) bool {
	action := strings.TrimRight(strings.TrimLeft(path, "/"), "/")

	if !strings.Contains(action, "git-") &&
		!strings.Contains(action, "info/") &&
		!strings.Contains(action, "HEAD") &&
		!strings.Contains(action, "objects/") {
		return false
	}

	return true
}

func (rt router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cleanPath := strings.TrimRight(strings.TrimLeft(r.URL.Path, "/"), "/")

	args := strings.Split(cleanPath, "/")

	uname := args[0]
	repo := cleanRepoName(args[1])

	if !IsGitRequest(cleanPath) {
		http.Error(w, "not found", http.StatusNotFound)
	}

	authHead := r.Header.Get("Authorization")
	if len(authHead) == 0 {
		w.Header().Set("WWW-Authenticate", "Basic realm=\".\"")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)

		return
	}

	auths := strings.Fields(authHead)
	if len(auths) != 2 || auths[0] != "Basic" {
		fmt.Println("middle")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	authUsername, authPassword, err := basicAuthDecode(auths[1])
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	shouldSend, code := rt.check(RequestInfo{
		RepoOwner: uname,
		RepoName:  repo,
		Username:  authUsername,
		Password:  authPassword,
	})

	if !shouldSend {
		http.Error(w, "", code)
		return
	}

	serve(&gitContext{
		w: w,
		r: r,
	})
}

func cleanRepoName(name string) string {
	var clean string

	clean = strings.TrimSuffix(name, ".git")
	clean = strings.TrimSuffix(clean, ".wiki")

	return clean
}

func getGitRepoPath(dir string) (string, error) {
	if !strings.HasSuffix(dir, ".git") {
		dir += ".git"
	}

	filename := path.Join(settings.RepoRootPath, dir)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return "", err
	}

	return filename, nil
}

func serve(ctx *gitContext) {
	for _, route := range routes {
		reqPath := strings.ToLower(ctx.r.URL.Path)
		m := route.GetMatches(reqPath)
		if m == nil {
			continue
		}

		if !route.IsMethod(ctx.r.Method) {
			ctx.NotFound()
			return
		}

		file := strings.TrimPrefix(reqPath, m[1]+"/")
		dir, err := getGitRepoPath(m[1])
		if err != nil {
			ctx.NotFound()
			return
		}

		route.handler(serviceHandler{
			w:    ctx.w,
			r:    ctx.r,
			file: file,
			dir:  dir,
		})
		return
	}

	ctx.NotFound()
}

func basicAuthDecode(encoded string) (string, string, error) {
	s, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", "", err
	}

	auth := strings.SplitN(string(s), ":", 2)
	return auth[0], auth[1], nil
}

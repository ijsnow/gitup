## Serve git over HTTP

### Intallation

```
$ go get github.com/ijsnow/gittp
```

### Documentation

#### `func NewHandler(rootRepoPath string, check func(RequestInfo) (isPass bool, code int)) http.Handler`

- `rootRepoPath` - Path to the directory where repositories are stored

- `check` - A function to check if we should handle the request. If `isPass` == true, we serve, else respond with the http status `code`.

#### `type RequestInfo struct`

```
type RequestInfo struct {
	RepoOwner string // Owner's username for requested repo
	RepoName  string // Requested repo name
	Username string // Username from git-credentials
	Password string // Password from git-credentials
}
```

#### `func IsGitRequest(uri string) bool`

- `uri` - The uri from the request to be checked

### Usage Example

```
package main

import (
	"fmt"
	"log"
	"net/http"
	"os/user"

	"github.com/ijsnow/gittp"
)

const (
  myGitUsername = "isaac"
  myGitPassord = "supersecret"
)

func check(r gittp.RequestInfo) (bool, int) {
  if r.Username == "im-definitely-not-a-hacker" {
    return false, http.StatusUnauthorized
  }

  if r.RepoName != r.Username {
    return false, http.StatusNotFound
  }

  if r.Username != myGitUsername || r.Password != myGitPassord {
    return false, http.StatusNotFound
  }

  return true, 0
}

func main() {
  usr, err := user.Current()
  if err != nil {
    log.Fatal(err)
  }
  repoDir := fmt.Sprintf("%s/my/git/repos", usr.HomeDir)

  githandler := gittp.NewHandler(repoDir, check)

  http.ListenAndServe(":3000", githandler)
}
```

### Special thanks

Thanks to those who built [gogs](https://github.com/gogits/gogs) for creating a great project and open sourcing for many others to use freely.
This code is a modified version of the
[module](https://github.com/gogits/gogs/blob/master/routers/repo/http.go)
that serves git over http in the gogs project.

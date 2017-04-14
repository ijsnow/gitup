package router

import (
	"net/http"
	"regexp"

	"github.com/ijsnow/gittp"
)

type rootHandler struct {
	git http.Handler
	web http.Handler
	api http.Handler
}

var isAPIRoute = regexp.MustCompile("^/api")

func (r rootHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	uri := req.URL.RequestURI()

	if gittp.IsGitRequest(uri) {
		r.git.ServeHTTP(w, req)
	} else if isAPIRoute.MatchString(uri) {
		r.api.ServeHTTP(w, req)
	} else {
		r.web.ServeHTTP(w, req)
	}
}

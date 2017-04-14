package gittp

import "net/http"

var settings = struct {
	RepoRootPath string
}{}

// SetRepoRootPath sets the root path for the directory storing repositories
func SetRepoRootPath(path string) {
	settings.RepoRootPath = path
}

// NewHandler returns a new handler to serve git repos
func NewHandler(path string, check func(RequestInfo) (bool, int)) http.Handler {
	settings.RepoRootPath = path

	return router{check}
}

// NewRouter returns a new mux
func NewRouter(path string, check func(RequestInfo) (bool, int)) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", NewHandler(path, check))

	return mux
}

package router

import (
	"io"
	"net/http"
)

type webHandler struct{}

func (r webHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Gitup Web Server")
}

func getWebHandler() http.Handler {
	return webHandler{}
}

package http

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Handler is used to create handlers for middlewares.
// It implements the http.Handler interface
type Handler struct {
	handlerFunc HandlerFunc
	middlewares []Middleware
}

// HandlerFunc is what handles the requests.
// It returns the response code and the data we are sending
type HandlerFunc func(http.ResponseWriter, *http.Request) (int, interface{})

// Middleware is the type for middlewares
type Middleware func(*http.Request) (*http.Request, int, interface{})

// NewHandler creates a new Handler
func NewHandler(handle HandlerFunc) Handler {
	return Handler{handle, make([]Middleware, 0)}
}

func (h Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var code int
	var data interface{}
	var err interface{}

	for _, mw := range h.middlewares {
		req, code, err = mw(req)

		if err != nil {
			break
		}
	}

	if err == nil {
		code, data = h.handlerFunc(res, req)
	} else {
		data = err
	}

	writeResponse(res, req, code, data)
}

// UseMiddlewares applies one or more middlewares for a HandlerFunc
func (h Handler) UseMiddlewares(middlewares ...Middleware) Handler {
	h.middlewares = middlewares

	return h
}

// writeResponse marshals data to a json struct and sends appropriate headers to w
func writeResponse(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	b, err := json.Marshal(data)

	if err != nil {
		log.Print(fmt.Sprintf("Error while encoding JSON: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Internal server error")
	} else {
		w.WriteHeader(code)
		io.WriteString(w, string(b))
	}
}

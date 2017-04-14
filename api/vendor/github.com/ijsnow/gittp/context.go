package gittp

import "net/http"

type gitContext struct {
	r *http.Request
	w http.ResponseWriter
}

// NotFound sends a not found response
func (h *gitContext) NotFound() {
	http.Error(h.w, "not found", http.StatusNotFound)
}

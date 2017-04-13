package controllers

import (
	"net/http"
)

// Git is the controller for git repo routes
type Git struct {
	baseController
}

// NewGitController returns a new Git controller
func NewGitController() *Git {
	return &Git{newBaseController()}
}

// GetRepo is the handler for serving git repos
func (c *Git) GetRepo(w http.ResponseWriter, r *http.Request) (int, interface{}) {
	return http.StatusCreated, nil
}

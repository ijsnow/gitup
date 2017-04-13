package controllers

import (
	"net/http"
)

// Core is the controller for authentication routes
type Core struct {
	baseController
}

// NewCoreController returns a new Core controller
func NewCoreController() *Core {
	return &Core{newBaseController()}
}

// Status is the handler for returning
func (c *Core) Status(w http.ResponseWriter, r *http.Request) (int, interface{}) {
	return http.StatusOK, nil
}

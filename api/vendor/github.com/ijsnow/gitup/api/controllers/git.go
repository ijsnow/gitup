package controllers

import (
	"net/http"

	"github.com/ijsnow/gitup/api/services/accounts"
	"github.com/ijsnow/gitup/api/services/repos"
	"github.com/ijsnow/gitup/types"

	"github.com/ijsnow/gittp"
)

// Git is the controller for git repo routes
type Git struct {
	baseController
}

// NewGitController returns a new Git controller
func NewGitController() *Git {
	return &Git{newBaseController()}
}

// ServeRepo is the handler for ensuring that we can serve a repo
func (c *Git) ServeRepo(r gittp.RequestInfo) (bool, int) {
	possibleUser := types.LoginUser{
		Uname:    r.Username,
		Password: r.Password,
	}

	_, ok := possibleUser.ValidateLogin()
	if !ok {
		return false, http.StatusBadRequest
	}

	user := types.User{}
	if accounts.GetUserByLogin(&possibleUser, &user) != nil {
		return false, http.StatusUnauthorized
	}

	repo := types.Repo{
		Name:    r.RepoName,
		OwnerID: user.ID,
	}

	err := repos.GetRepo(&repo)
	if err != nil {
		return false, http.StatusNotFound
	}

	// TODO: Also check for collaborators
	if repo.IsPrivate && repo.OwnerID != user.ID {
		return false, http.StatusUnauthorized
	}

	return true, http.StatusOK
}

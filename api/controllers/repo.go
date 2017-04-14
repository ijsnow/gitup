package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gitup.io/isaac/gitup/api/services/middlewares"
	"gitup.io/isaac/gitup/api/services/repos"
	"gitup.io/isaac/gitup/types"
	httputils "gitup.io/isaac/gitup/utils/http"
)

// Repo is the controller for managing remote repos
type Repo struct {
	baseController
}

// NewRepoController returns a new Repo controller
func NewRepoController() *Repo {
	return &Repo{newBaseController()}
}

// CreateRepo is the handler for creating a repo
func (c *Repo) CreateRepo(w http.ResponseWriter, req *http.Request) (int, interface{}) {
	sctx := req.Context().Value(middlewares.SessionKey)
	session, ok := sctx.(*types.Session)
	if !ok {
		fmt.Println("cant convert to session", sctx)
		return httputils.InternalError()
	}

	repo := types.Repo{}

	err := json.NewDecoder(req.Body).Decode(&repo)
	if err != nil {
		return httputils.InternalError()
	}

	repo.Uname = session.User.Uname
	repo.OwnerID = session.User.ID

	err = repos.CreateBareRepo(&repo)
	if err != nil {
		return http.StatusConflict, httputils.ErrorResponse([]string{"Repo already exists"})
	}

	return http.StatusOK, nil
}

// DeleteRepo is the handler for deleting a repo
func (c *Repo) DeleteRepo(w http.ResponseWriter, req *http.Request) (int, interface{}) {
	sctx := req.Context().Value(middlewares.SessionKey)
	session, ok := sctx.(*types.Session)
	if !ok {
		return httputils.InternalError()
	}

	query := req.URL.Query()

	repo := types.Repo{
		Name: query["name"][0],
	}

	repo.Uname = session.User.Uname

	err := repos.DeleteRepo(&repo)
	if err != nil {
		return http.StatusConflict, httputils.ErrorResponse([]string{"Repo does not exist"})
	}

	return http.StatusOK, nil
}

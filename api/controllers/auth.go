package controllers

import (
	"encoding/json"
	"net/http"

	"gitup.io/isaac/gitup/api/services/accounts"
	"gitup.io/isaac/gitup/api/services/middlewares"
	"gitup.io/isaac/gitup/api/services/repos"
	"gitup.io/isaac/gitup/api/services/sessions"
	"gitup.io/isaac/gitup/types"
	httputils "gitup.io/isaac/gitup/utils/http"
)

// Auth is the controller for authentication routes
type Auth struct {
	baseController
}

// NewAuthController returns a new Authentication controller
func NewAuthController() *Auth {
	return &Auth{newBaseController()}
}

// Login is the handler for signing up from a client
func (c *Auth) Login(w http.ResponseWriter, req *http.Request) (int, interface{}) {
	possibleUser := types.LoginUser{}

	err := json.NewDecoder(req.Body).Decode(&possibleUser)
	if err != nil {
		return httputils.InternalError()
	}

	errors, ok := possibleUser.ValidateLogin()
	if !ok {
		return http.StatusBadRequest, httputils.ErrorResponse(errors)
	}

	user := types.User{}
	// Validate with DB
	if accounts.GetUserByLogin(&possibleUser, &user) != nil {
		return httputils.NotFound()
	}

	device := req.Header.Get("Client-Device")

	err = sessions.CreateSession(&user, device)
	if err != nil {
		return httputils.InternalError()
	}

	return http.StatusCreated, user
}

// CreateUser creates a user on sign up
func (c *Auth) CreateUser(w http.ResponseWriter, req *http.Request) (int, interface{}) {
	possibleUser := types.LoginUser{}

	err := json.NewDecoder(req.Body).Decode(&possibleUser)
	if err != nil {
		return httputils.InternalError()
	}

	// Validate sign up values
	errors, ok := possibleUser.ValidateSignup()
	if !ok {
		return http.StatusBadRequest, httputils.ErrorResponse(errors)
	}

	user := types.User{
		Uname:    possibleUser.Uname,
		Email:    possibleUser.Email,
		Password: possibleUser.Password,
	}

	// Validate with DB
	err = accounts.ValidateUser(user)
	if err != nil {
		return http.StatusConflict, httputils.ErrorResponse([]string{err.Error()})
	}

	err = accounts.CreateUser(&user)
	if err != nil {
		return httputils.InternalError()
	}

	device := req.Header.Get("Client-Device")

	err = sessions.CreateSession(&user, device)
	if err != nil {
		return httputils.InternalError()
	}

	err = repos.ProvisionRepos(&user)
	if err != nil {
		return httputils.InternalError()
	}

	return http.StatusCreated, map[string]interface{}{"token": user.Token}
}

// Logout logs the user out by destroying the session
func (c *Auth) Logout(w http.ResponseWriter, req *http.Request) (int, interface{}) {
	sctx := req.Context().Value(middlewares.SessionKey)
	session, ok := sctx.(*types.Session)
	if !ok {
		return httputils.InternalError()
	}

	accounts.Logout(session)

	return http.StatusOK, nil
}

package api

import (
	"gitup.io/isaac/gitup/types"
)

// AuthResp is the response type for a authentication endpoint
type AuthResp struct {
	responseBase
	types.User
}

// Login attempts to log the user in
func Login(user *types.User) (bool, AuthResp) {
	resp := AuthResp{}

	p := params{
		"username": user.Uname,
		"password": user.Password,
	}

	err := post("auth/login", p, &resp)
	if err != nil {
		return false, resp
	}

	user.Token = resp.Token
	user.Email = resp.Email

	return resp.Success, resp
}

// Signup attempts to sign the user up
func Signup(user *types.User) (bool, AuthResp) {
	resp := AuthResp{}

	p := params{
		"username": user.Uname,
		"email":    user.Email,
		"password": user.Password,
	}

	err := post("auth/signup", p, &resp)
	if err != nil {
		return false, resp
	}

	user.Token = resp.Token
	user.Email = resp.Email

	return resp.Success, resp
}

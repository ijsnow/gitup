package api

import (
	"fmt"

	"github.com/ijsnow/gitup/types"
)

// AuthResp is the response type for a authentication endpoint
type AuthResp struct {
	responseBase
	types.User
}

// Login attempts to log the user in
func Login(user *types.User) (bool, AuthResp) {
	resp := AuthResp{}

	fmt.Println("api")

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

// Logout destroys the session on the server
func Logout() bool {
	resp := AuthResp{}
	authDelete("auth/logout", nil, &resp)

	return resp.Success
}

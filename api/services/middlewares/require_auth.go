package middlewares

import (
	"context"
	"net/http"

	"gitup.io/isaac/gitup/datastore"
	"gitup.io/isaac/gitup/types"
)

type key string

const (
	// SessionKey is the key where we will store users
	SessionKey key = "session-key"
)

// RequireAuth is a middleware for requiring authentication
func RequireAuth(req *http.Request) (*http.Request, int, interface{}) {
	token := req.Header.Get("Authorization")
	session := &types.Session{}

	err := datastore.Store.Sessions.GetSessionByToken(token, session)
	if err != nil {
		return req, http.StatusUnauthorized, "Unauthorized"
	}

	var user types.User

	datastore.Store.Users.GetUserByID(session.UID, &user)
	if err != nil {
		return req, http.StatusInternalServerError, nil
	}

	session.User = &user

	ctx := req.Context()
	ctx = context.WithValue(ctx, SessionKey, session)

	return req.WithContext(ctx), 0, nil
}

package sessions

import (
	"errors"

	"github.com/ijsnow/gitup/datastore"
	"github.com/ijsnow/gitup/types"
)

// CreateSession creates a new session for the user
func CreateSession(user *types.User, device string) error {
	ds := datastore.Store

	token, err := ds.Sessions.NewSession(user.ID, device)
	if err != nil {
		return errors.New("Unable to create session")
	}

	user.Token = token

	return nil
}

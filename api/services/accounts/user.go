package accounts

import (
	"errors"

	"gitup.io/isaac/gitup/datastore"
	"gitup.io/isaac/gitup/services/authentication"
	"gitup.io/isaac/gitup/types"
)

// ValidateUser validate the user
func ValidateUser(user types.User) error {
	if datastore.Store.Users.UserExistsByUname(user.Uname) {
		return errors.New("Username is associated with another account")
	}

	return nil
}

// CreateUser creates a new user and returns it's representation in the DB
func CreateUser(user *types.User) error {
	hash, err := authentication.PasswordHash(user.Password)
	if err != nil {
		return errors.New("Could not create user")
	}

	user.PasswordHash = hash

	id, err := datastore.Store.Users.CreateUser(user)
	if err != nil {
		return errors.New("Could not create user in DB")
	}

	user.ID = id

	return nil
}

// GetUserByUname gets a user by uname
func GetUserByUname(user *types.User) error {
	return datastore.Store.Users.GetUserByUname(user)
}

// GetUserByLogin gets a user by login attempt
func GetUserByLogin(login *types.LoginUser, user *types.User) error {
	return datastore.Store.Users.GetUserByLogin(login, user)
}

// AuthenticateUser checks the users credentials
func AuthenticateUser(user *types.User, login *types.LoginUser) bool {
	return authentication.Authenticate(user, login)
}

// Logout destroys the session
func Logout(session *types.Session) error {
	return datastore.Store.Sessions.DestroySession(session.Token)
}

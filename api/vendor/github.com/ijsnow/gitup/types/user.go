package types

import (
	"time"

	"github.com/ijsnow/gitup/services/validate"

	"github.com/asaskevich/govalidator"
)

const (
	// StatusActive user type
	StatusActive = iota
	// StatusDisabled user type
	StatusDisabled
	// Archived user type
	Archived
	// StatusSuspended user type
	StatusSuspended
)

// User is the user type. It is safe to be sent over the wire
type User struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	Uname        string    `json:"username"`
	Status       int       `json:"status"`
	Token        string    `json:"token"`
	Password     string    `json:"_"`
	PasswordHash string    `json:"_"`
	CreatedAt    time.Time `json:"createdAt"`
}

// DBUser is the type of user that is stored in a db. This should not be sent over the wire
type DBUser struct {
	ID           int
	Email        string
	Uname        string
	Status       int
	PasswordHash string
	CreatedAt    time.Time
}

// LoginUser is the type that we use to marshall/unmarshall auth attempts
type LoginUser struct {
	Email    string `json:"email"`
	Uname    string `json:"username"`
	Password string `json:"password"`
}

// NewUser creates a new user struct
func NewUser(email, username string) User {
	return User{
		Email:  email,
		Uname:  username,
		Status: StatusActive,
	}
}

// ValidateSignup is used for validating a new user when signing up
func (u *LoginUser) ValidateSignup() ([]string, bool) {
	errors := make([]string, 0)

	if u.Email == "" {
		errors = append(errors, "Email is required")
	} else if !govalidator.IsEmail(u.Email) {
		errors = append(errors, "Email is invalid")
	}

	if u.Uname == "" {
		errors = append(errors, "Username is required")
	} else if !validate.Uname(u.Uname) {
		errors = append(errors, "Username is invalid")
	}

	if u.Password == "" {
		errors = append(errors, "Password is required")
	} else if !validate.Password(u.Password) {
		errors = append(errors, "Password is invalid")
	}

	return errors, len(errors) == 0
}

// ValidateLogin is used for validating a user when logging in
func (u *LoginUser) ValidateLogin() ([]string, bool) {
	errors := make([]string, 0)

	if u.Uname == "" {
		errors = append(errors, "Username is required")
	}

	if u.Password == "" {
		errors = append(errors, "Password is required")
	}

	return errors, len(errors) == 0
}

// GetStatusString gets the users status as a string
func (u *User) GetStatusString() string {
	switch u.Status {
	case StatusActive:
		return "active"
	case StatusDisabled:
		return "disabled"
	case Archived:
		return "archived"
	default:
		return "suspended"
	}
}

// ToDBUser convers a User to a DBUser
func (u *User) ToDBUser() *DBUser {
	return &DBUser{
		ID:           u.ID,
		Email:        u.Email,
		Uname:        u.Uname,
		Status:       u.Status,
		PasswordHash: u.PasswordHash,
		CreatedAt:    u.CreatedAt,
	}
}

// ToUser convers a User to a DBUser
func (u *DBUser) ToUser(us *User) {
	us.ID = u.ID
	us.Email = u.Email
	us.Uname = u.Uname
	us.Status = u.Status
	us.PasswordHash = u.PasswordHash
	us.CreatedAt = u.CreatedAt
}

package bolt

import (
	"errors"
	"fmt"

	boltdb "github.com/boltdb/bolt"
	"gitup.io/isaac/gitup/services/authentication"
	"gitup.io/isaac/gitup/types"
)

// Users is the structure for managing the users in the db
type Users struct {
	db *boltdb.DB
}

// NewUsers initiallizes the users portion of the db
func NewUsers(db *boltdb.DB) *Users {
	db.Update(func(tx *boltdb.Tx) error {
		_, err := tx.CreateBucketIfNotExists(keys.user)
		if err != nil {
			return err
		}
		return nil
	})

	return &Users{
		db: db,
	}
}

// CreateUser creates a user object
func (c *Users) CreateUser(newUser *types.User) (int, error) {
	err := c.db.Update(func(tx *boltdb.Tx) error {
		b := tx.Bucket(keys.user)

		id, _ := b.NextSequence()
		newUser.ID = int(id)

		u := newUser.ToDBUser()

		enc, err := encode(u)
		if err != nil {
			return err
		}

		err = b.Put(itob(newUser.ID), enc)

		return err
	})

	return newUser.ID, err
}

// GetUserByUname gets a user by username
func (c *Users) GetUserByUname(user *types.User) error {
	return c.db.View(func(tx *boltdb.Tx) error {
		var err error
		c := tx.Bucket(keys.user).Cursor()
		u := &types.DBUser{}
		isExist := false

		for k, v := c.First(); k != nil; k, v = c.Next() {
			err = decode(v, u)

			if user.Uname == u.Uname {
				u.ToUser(user)
				isExist = true
				fmt.Println(2, u.PasswordHash)

				break
			}

			if err != nil {
				return err
			}
		}

		if !isExist {
			return errors.New("Not found")
		}

		return err
	})
}

// UserExistsByUname sees if a user exists
func (c *Users) UserExistsByUname(uname string) bool {
	if err := c.GetUserByUname(&types.User{Uname: uname}); err == nil {
		return true
	}

	return false
}

// GetUserByLogin gets a user by login request. It checks the validity of the password and email
func (c *Users) GetUserByLogin(login *types.LoginUser, user *types.User) error {
	user.Uname = login.Uname
	err := c.GetUserByUname(user)
	if err != nil {
		return err
	}

	fmt.Println("Found user", user.Uname, user.Email, 1, user.PasswordHash)

	if !authentication.Authenticate(user, login) {
		return errors.New("Not found")
	}

	return nil
}

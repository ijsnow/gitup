package bolt

import (
	"fmt"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ijsnow/gitup/api/config"
	"github.com/ijsnow/gitup/services/authentication"
	"github.com/ijsnow/gitup/types"

	boltdb "github.com/boltdb/bolt"
)

// Sessions is the structure for managing the sessions in the db
type Sessions struct {
	db *boltdb.DB
}

// NewSessions initiallizes the users portion of the db
func NewSessions(db *boltdb.DB) *Sessions {
	db.Update(func(tx *boltdb.Tx) error {
		_, err := tx.CreateBucketIfNotExists(keys.session)
		if err != nil {
			return err
		}
		return nil
	})

	return &Sessions{
		db: db,
	}
}

// NewSession creates a new session in the database
func (c *Sessions) NewSession(userID int, deviceType string) (string, error) {
	token, err := authentication.GenerateJwt(strconv.Itoa(userID))
	if err != nil {
		fmt.Println(1, err)
		return "", err
	}

	err = c.db.Update(func(tx *boltdb.Tx) error {
		b := tx.Bucket(keys.session)

		id, _ := b.NextSequence()

		t := types.NewSession(userID, token, deviceType)
		t.ID = int(id)

		var enc []byte
		enc, err = encode(t)
		if err != nil {
			return err
		}

		err = b.Put([]byte(token), enc)

		return err
	})

	return token, err
}

// GetSessionByToken gets a session by a token
func (c *Sessions) GetSessionByToken(token string, session *types.Session) error {
	return c.db.View(func(tx *boltdb.Tx) error {
		b := tx.Bucket(keys.session)
		k := []byte(token)

		return decode(b.Get(k), session)
	})
}

// DestroySession destroys a session by a given token
func (c *Sessions) DestroySession(token string) error {
	return c.db.Update(func(tx *boltdb.Tx) error {
		return tx.Bucket(keys.session).Delete([]byte(token))
	})
}

// jwt.Token https://godoc.org/github.com/dgrijalva/jwt-go#Token
// type Token struct {
//     Raw       string                 // The raw token.  Populated when you Parse a token
//     Method    SigningMethod          // The signing method used or to be used
//     Header    map[string]interface{} // The first segment of the token
//     Claims    Claims                 // The second segment of the token
//     Signature string                 // The third segment of the token.  Populated when you Parse a token
//     Valid     bool                   // Is the token valid?  Populated when you Parse/Verify a token
// }

func isTokenValid(s *types.Session) bool {
	token, err := parseToken(s)
	if err != nil || !token.Valid {
		return false
	}

	return true
}

func parseToken(s *types.Session) (*jwt.Token, error) {
	return jwt.Parse(s.Token, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.App.Security.Secret), nil
	})
}

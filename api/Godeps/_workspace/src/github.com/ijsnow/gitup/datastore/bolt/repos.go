package bolt

import (
	"errors"

	boltdb "github.com/boltdb/bolt"
	"github.com/ijsnow/gitup/types"
)

// Repos is the structure for managing the repos in the db
type Repos struct {
	db *boltdb.DB
}

// NewRepos initiallizes the users portion of the db
func NewRepos(db *boltdb.DB) *Repos {
	db.Update(func(tx *boltdb.Tx) error {
		_, err := tx.CreateBucketIfNotExists(keys.repo)
		if err != nil {
			return err
		}
		return nil
	})

	return &Repos{
		db: db,
	}
}

// CreateUserBucket creates a bucket to store a user's repos in
func (c *Repos) CreateUserBucket(uID int) error {
	return c.db.Update(func(tx *boltdb.Tx) error {
		b := tx.Bucket(keys.repo)

		_, err := b.CreateBucket(itob(uID))

		return err
	})
}

// CreateRepo creates a repo object
func (c *Repos) CreateRepo(r *types.Repo) error {
	// Start the transaction.
	tx, err := c.db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	root := tx.Bucket(keys.repo)

	b, err := root.CreateBucketIfNotExists(itob(r.OwnerID))
	if err != nil {
		return err
	}

	// Generate an ID for the new user.
	rID, err := b.NextSequence()
	if err != nil {
		return err
	}
	r.ID = int(rID)

	enc, err := encode(r)
	if err != nil {
		return err
	}

	err = b.Put(itob(r.ID), enc)
	if err != nil {
		return err
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

// GetRepoByName gets a repo by user name
func (c *Repos) GetRepoByName(repo *types.Repo) error {
	return c.db.View(func(tx *boltdb.Tx) error {
		root := tx.Bucket(keys.repo)
		b := root.Bucket(itob(repo.OwnerID))
		if b == nil {
			return errors.New("No repos for this user")
		}

		var err error
		c := b.Cursor()
		r := &types.Repo{}
		isExist := false

		for k, v := c.First(); k != nil; k, v = c.Next() {
			err = decode(v, r)

			if repo.Name == r.Name {
				r.Copy(repo)
				isExist = true

				break
			}

			if err != nil {
				return err
			}
		}

		if !isExist {
			return ErrNotFound
		}

		return err
	})
}

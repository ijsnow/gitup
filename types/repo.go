package types

import (
	"fmt"
	"time"
)

// Repo is a data structure to represent a git repo
type Repo struct {
	ID        int
	Name      string    `json:"name"`
	Uname     string    `json:"_"`
	OwnerID   int       `json:"ownerID"`
	IsPrivate bool      `json:"isPrivate"`
	CreatedAt time.Time `json:"createdAt"`
}

// PathName builds the path name
func (r Repo) PathName() string {
	return fmt.Sprintf("%s/%s.git", r.Uname, r.Name)
}

// Copy assigns all the values of this Repo to another Repo
func (r Repo) Copy(repo *Repo) {
	repo.ID = r.ID
	repo.Name = r.Name
	repo.Uname = r.Uname
	repo.OwnerID = r.OwnerID
	repo.IsPrivate = r.IsPrivate
	repo.CreatedAt = r.CreatedAt
}

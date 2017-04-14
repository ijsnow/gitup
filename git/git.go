package git

import (
	lgit "gopkg.in/libgit2/git2go.v25"
)

// Repository is a repository type
type Repository lgit.Repository

// CreateBareRepo creates a repo at a given path
func CreateBareRepo(path string) (*lgit.Repository, error) {
	return lgit.InitRepository(path, true)
}

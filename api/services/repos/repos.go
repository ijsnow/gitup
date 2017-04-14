package repos

import (
	"errors"
	"fmt"
	"os"

	"github.com/ijsnow/gitup/api/config"
	"github.com/ijsnow/gitup/datastore"
	"github.com/ijsnow/gitup/git"
	"github.com/ijsnow/gitup/types"
	"github.com/ijsnow/gitup/utils/fspaths"
)

var repoDir string

// CreateBareRepo creates a repo for the current user
func CreateBareRepo(r *types.Repo) error {
	err := datastore.Store.Repos.CreateRepo(r)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s/%s", repoDir, r.PathName())
	if _, err = os.Stat(path); err == nil {
		return errors.New("Repo already exists")
	}

	_, err = git.CreateBareRepo(path)

	return err
}

// DeleteRepo deletes a repo
func DeleteRepo(r *types.Repo) error {
	path := fmt.Sprintf("%s/%s", repoDir, r.PathName())
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return errors.New("Repo does not exist")
	}

	os.RemoveAll(path)

	return nil
}

// ProvisionRepos creates the directories needed for repos when a user signs up
// as well as creates the user's repo info in the DB
func ProvisionRepos(user *types.User) error {
	err := datastore.Store.Repos.CreateUserBucket(user.ID)
	if err != nil {
		return err
	}

	err = os.MkdirAll(fmt.Sprintf("%s/%s", repoDir, user.Uname), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// GetRepo gets the repo with name and user info
func GetRepo(repo *types.Repo) error {
	return datastore.Store.Repos.GetRepoByName(repo)
}

// GetRepoDir gets the fully built repo directory
func GetRepoDir() string {
	return repoDir
}

func init() {
	path := fspaths.ExpandTildePath(config.App.RepoDir)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}

	repoDir = path
}

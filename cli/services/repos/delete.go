package repos

import (
	"os"
	"strings"

	"gitup.io/isaac/gitup/cli/api"
	"gitup.io/isaac/gitup/cli/config"
	"gitup.io/isaac/gitup/iocli"
	"gitup.io/isaac/gitup/types"
)

// DeleteRemoteRepo creates a remote repository
func DeleteRemoteRepo() error {
	dir, _ := os.Getwd()
	dir = strings.Replace(dir, " ", "\\ ", -1)
	dir = dir[strings.LastIndex(dir, "/")+1:]

	repoName := dir

	repo := types.Repo{Name: repoName, Uname: config.Uname}
	iocli.Error("Are you sure you want to the remote repo at %s/%s?", config.Host, repo.PathName())
	inp := iocli.PromptString("Enter the name of this repo to confirm (%s)", repoName)

	if inp.Response != repoName {
		iocli.Info("Not deleting. Probably a safe decision")
		return nil
	}

	iocli.Info("Deleting remote repo at %s/%s", config.Host, repo.PathName())
	success, resp := api.DeleteRepo(&repo)
	if success {
		iocli.Success("Repo %s was successfully deleted", repoName)
	} else {
		iocli.Error("Unable to delete remote repo")
		iocli.Errors(resp.Errors)
	}

	return nil
}

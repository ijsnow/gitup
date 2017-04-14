package repos

import (
	"os"
	"strings"

	"gitup.io/isaac/gitup/cli/api"
	"gitup.io/isaac/gitup/cli/config"
	"gitup.io/isaac/gitup/iocli"
	"gitup.io/isaac/gitup/services/validate"
	"gitup.io/isaac/gitup/types"
)

// CreateRemoteRepo creates a remote repository
func CreateRemoteRepo(name string) error {
	repoName := name
	if name == "" {
		dir, _ := os.Getwd()
		dir = strings.Replace(dir, " ", "\\ ", -1)
		dir = dir[strings.LastIndex(dir, "/")+1:]

		repoName = dir
	}

	repo := types.Repo{Name: repoName, Uname: config.Uname}

	inp := iocli.PromptRune("Create a remote repo at %s/%s? [Y/n]", config.Host, repo.PathName())
	if inp.IsNo() {
		inp = iocli.PromptRepoName("Enter the repo name you would like(or press ^c to quit)")
		repo.Name = inp.Response
	}

	for !validate.RepoName(repoName) {
		iocli.Error("Oops! The repo name you entered was was invalid.")
		inp = iocli.PromptRepoName("Enter a repo name")
		repo.Name = inp.Response
	}

	inp = iocli.PromptRune("Is this a private repo? [Y/n]")
	repo.IsPrivate = inp.IsYes()

	iocli.Info("Creating remote repo at %s/%s", config.Host, repo.PathName())
	success, resp := api.CreateRemoteRepo(&repo)
	if success {
		iocli.Success("Repo %s was successfully created", repoName)
		iocli.Info("To add your project to the new remote repo, run the following commands")
		iocli.Info("  git init")
		iocli.Info("  git add .")
		iocli.Info(`  git commit -m "initial commit"`)
		iocli.Info("  git remote add origin %s/%s", config.Host, repo.PathName())
		iocli.Info("  git push -u origin master")
	} else {
		iocli.Error("Unable to create remote repo")
		iocli.Errors(resp.Errors)
	}

	return nil
}

package repo

import (
	"github.com/urfave/cli"
	"gitup.io/isaac/gitup/cli/config"
	"gitup.io/isaac/gitup/cli/services/repos"
	"gitup.io/isaac/gitup/cli/utils"
)

func createRepoAction(c *cli.Context) error {
	repoName := c.Args().First()

	return repos.CreateRemoteRepo(repoName)
}

// Create attempts to create remote repo
var Create = cli.Command{
	Name:   "create",
	Usage:  "Create a remote repo",
	Action: utils.CreateAction(config.RequireAuth(createRepoAction)),
}

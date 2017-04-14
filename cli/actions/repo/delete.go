package repo

import (
	"github.com/urfave/cli"
	"gitup.io/isaac/gitup/cli/config"
	"gitup.io/isaac/gitup/cli/services/repos"
	"gitup.io/isaac/gitup/cli/utils"
)

func deleteRepoAction(c *cli.Context) error {
	return repos.DeleteRemoteRepo()
}

// Delete attempts to create remote repo
var Delete = cli.Command{
	Name:   "delete",
	Usage:  "Delete a remote repo",
	Action: utils.CreateAction(config.RequireAuth(deleteRepoAction)),
}

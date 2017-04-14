package repo

import (
	"github.com/urfave/cli"
	"github.com/ijsnow/gitup/cli/config"
	"github.com/ijsnow/gitup/cli/services/repos"
	"github.com/ijsnow/gitup/cli/utils"
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

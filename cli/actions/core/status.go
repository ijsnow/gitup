package core

import (
	"github.com/urfave/cli"
	"github.com/ijsnow/gitup/cli/api"
	"github.com/ijsnow/gitup/cli/config"
	"github.com/ijsnow/gitup/cli/utils"
	"github.com/ijsnow/gitup/iocli"
)

func statusAction(c *cli.Context) error {
	if status, _ := api.Status(); status {
		iocli.Success("gitup API is up and running!")
	} else {
		iocli.Error("gitup API is down :(")
	}

	return nil
}

// Status checks the status of the api
var Status = cli.Command{
	Name:   "status",
	Usage:  "Check the status of the gitup API",
	Action: utils.CreateAction(config.RequireHost(statusAction)),
}

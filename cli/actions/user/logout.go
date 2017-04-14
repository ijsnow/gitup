package user

import (
	"github.com/urfave/cli"
	"github.com/ijsnow/gitup/cli/config"
	"github.com/ijsnow/gitup/cli/services/auth"
	"github.com/ijsnow/gitup/cli/utils"
)

func logoutAction(c *cli.Context) error {
	return auth.Logout()
}

// Logout logs the user out
var Logout = cli.Command{
	Name:   "logout",
	Usage:  "Log the current user out",
	Action: utils.CreateAction(config.RequireHost(logoutAction)),
}

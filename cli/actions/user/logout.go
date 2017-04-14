package user

import (
	"github.com/urfave/cli"
	"gitup.io/isaac/gitup/cli/config"
	"gitup.io/isaac/gitup/cli/services/auth"
	"gitup.io/isaac/gitup/cli/utils"
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

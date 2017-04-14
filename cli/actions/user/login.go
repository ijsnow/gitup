package user

import (
	"github.com/urfave/cli"
	"github.com/ijsnow/gitup/cli/config"
	"github.com/ijsnow/gitup/cli/services/auth"
	"github.com/ijsnow/gitup/cli/utils"
)

func loginAction(c *cli.Context) error {
	return auth.Login()
}

// Login attempts to log in the user or create an account for new users
var Login = cli.Command{
	Name:   "login",
	Usage:  "Log in with or create a user account",
	Action: utils.CreateAction(config.RequireHost(loginAction)),
}

package user

import (
	"github.com/urfave/cli"
	"github.com/ijsnow/gitup/cli/config"
	"github.com/ijsnow/gitup/cli/services/auth"
	"github.com/ijsnow/gitup/cli/utils"
	"github.com/ijsnow/gitup/types"
)

func signupAction(c *cli.Context) error {
	return auth.Signup(&types.User{})
}

// Signup attempts to create an account for new users
var Signup = cli.Command{
	Name:   "signup",
	Usage:  "Create a user account",
	Action: utils.CreateAction(config.RequireHost(signupAction)),
}

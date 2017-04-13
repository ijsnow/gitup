package user

import (
	"github.com/urfave/cli"
	"gitup.io/isaac/gitup/cli/config"
	"gitup.io/isaac/gitup/cli/services/auth"
	"gitup.io/isaac/gitup/cli/utils"
	"gitup.io/isaac/gitup/types"
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

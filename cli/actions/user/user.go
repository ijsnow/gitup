package user

import (
	"github.com/urfave/cli"
	"github.com/ijsnow/gitup/cli/config"
	"github.com/ijsnow/gitup/cli/utils"
	"github.com/ijsnow/gitup/iocli"
)

func currentUserAction(c *cli.Context) error {
	if config.Uname != "" {
		iocli.Success("User %s is logged in", config.Uname)
	} else {
		iocli.Error("You are currently not logged in")
		iocli.Info("Run `gitup login`")
	}

	return nil
}

// User is the cli command that does all actions dealing with the user account
var User = cli.Command{
	Name:      "user",
	Usage:     "Log in with or create a gitup.io account",
	ArgsUsage: "<username>",
	Action:    utils.CreateAction(config.RequireConfig(currentUserAction)),
	Subcommands: []cli.Command{
		Signup,
		Login,
		Logout,
	},
}

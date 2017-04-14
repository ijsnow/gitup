package host

import (
	"github.com/urfave/cli"
	"github.com/ijsnow/gitup/cli/config"
	"github.com/ijsnow/gitup/cli/utils"
	"github.com/ijsnow/gitup/iocli"
)

func unsetHostAction(c *cli.Context) error {
	cfg := config.UserConfig{Host: ""}
	config.SaveConfig(cfg)

	return nil
}

// User is the cli command that does all actions dealing with the user account
var unset = cli.Command{
	Name:   "unset",
	Usage:  "Log in with or create a gitup.io account",
	Action: utils.CreateAction(config.RequireConfig(unsetHostAction)),
}

func hostAction(c *cli.Context) error {
	if config.Host != "" {
		iocli.Success("Your current host is %s", config.Host)
	}

	return nil
}

// Host tells you your current host
var Host = cli.Command{
	Name:   "host",
	Usage:  "Check your current host",
	Action: utils.CreateAction(config.RequireHost(hostAction)),
	Subcommands: []cli.Command{
		unset,
	},
}

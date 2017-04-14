package main

import (
	"os"

	"github.com/ijsnow/gitup/cli/actions/core"
	"github.com/ijsnow/gitup/cli/actions/host"
	"github.com/ijsnow/gitup/cli/actions/repo"
	"github.com/ijsnow/gitup/cli/actions/user"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{}

	app.Commands = []cli.Command{
		core.Status,
		user.User,
		user.Login,
		user.Signup,
		user.Logout,
		repo.Repo,
		host.Host,
	}

	app.Run(os.Args)
}

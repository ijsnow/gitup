package main

import (
	"os"

	"gitup.io/isaac/gitup/cli/actions/core"
	"gitup.io/isaac/gitup/cli/actions/host"
	"gitup.io/isaac/gitup/cli/actions/repo"
	"gitup.io/isaac/gitup/cli/actions/user"

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

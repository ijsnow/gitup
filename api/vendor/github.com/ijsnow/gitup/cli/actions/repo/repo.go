package repo

import (
	"github.com/urfave/cli"
)

// Repo is the cli command that does all actions dealing with repos
var Repo = cli.Command{
	Name: "repo",
	Subcommands: []cli.Command{
		Create,
		Delete,
	},
}

package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/ijsnow/gitup/iocli"
)

func unameAction(c *cli.Context) error {
	res := iocli.PromptUname("Enter Username")

	iocli.Success("Username entered: %s", res.Response)

	return nil
}

// Status checks the status of the api
var uname = cli.Command{
	Name:   "uname",
	Action: unameAction,
}

func emailAction(c *cli.Context) error {
	res := iocli.PromptEmail("Enter Email")

	iocli.Success("Email entered: %s", res.Response)

	return nil
}

// Status checks the status of the api
var email = cli.Command{
	Name:   "email",
	Action: emailAction,
}

func passwordAction(c *cli.Context) error {
	res := iocli.PromptPassword("Enter password")

	iocli.Success("Password entered: %s", res.Response)

	return nil
}

// Status checks the status of the api
var password = cli.Command{
	Name:   "password",
	Action: passwordAction,
}

func runeAction(c *cli.Context) error {
	res := iocli.PromptRune("Enter a rune")

	iocli.Success("Entered rune: %s", res.Response)

	return nil
}

// Status checks the status of the api
var runeA = cli.Command{
	Name:   "rune",
	Action: runeAction,
}

func hostAction(c *cli.Context) error {
	res := iocli.PromptHost("Enter a hostname")

	iocli.Success("Entered hostname: %s", res.Response)

	return nil
}

// Status checks the status of the api
var host = cli.Command{
	Name:   "host",
	Action: hostAction,
}

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{}

	app.Commands = []cli.Command{
		uname,
		email,
		password,
		runeA,
		host,
	}

	app.Run(os.Args)
}

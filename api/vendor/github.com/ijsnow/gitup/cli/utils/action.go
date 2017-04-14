package utils

import (
	"github.com/urfave/cli"
	"github.com/ijsnow/gitup/iocli"
)

// CreateAction wraps an action
func CreateAction(action func(*cli.Context) error) func(*cli.Context) error {
	return func(c *cli.Context) error {
		err := action(c)
		if err != nil {
			iocli.Error("%s", err)
		}

		return nil
	}
}

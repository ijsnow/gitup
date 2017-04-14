package auth

import (
	"github.com/ijsnow/gitup/cli/api"
	"github.com/ijsnow/gitup/cli/config"
	"github.com/ijsnow/gitup/iocli"
)

// Logout perges the user configs
func Logout() error {
	if config.Token == "" {
		iocli.Info("You are already logged out")
	}

	err := config.Logout()
	if err != nil {
		iocli.Error("Error logging out")
	} else {
		iocli.Success("You are now logged out")
	}

	api.Logout()

	return nil
}

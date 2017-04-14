package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"time"

	"github.com/ijsnow/gitup/iocli"

	"github.com/urfave/cli"
)

var (
	// Uname is the current computer's username
	Uname string
	// Email is the current computers email associated with the account
	Email string
	// Host is the remote host's URL or IP
	Host string

	// Token is the current user's token used for auth with gitup.io
	Token string

	configFilename = "config"
	keysFilename   = ".keys"
)

type config struct {
	Uname string `json:"username"`
	Email string `json:"email"`
	Host  string `json:"host"`
}

type keys struct {
	Key   string `json:"key"`
	Token string `json:"token"`
}

func getConfig() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	hdir := fromSystemPath(usr.HomeDir)

	configDir := toSystemPath(fmt.Sprintf("%s/.gitup/", hdir))
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		os.Mkdir(configDir, 0755)
	}

	configPath := toSystemPath(fmt.Sprintf("%s/%s", configDir, configFilename))
	if data, err := ioutil.ReadFile(configPath); err == nil {
		configData := config{}

		if err = json.Unmarshal(data, &configData); err == nil {
			Uname = configData.Uname
			Email = configData.Email
			Host = configData.Host
		}
	}

	keysPath := toSystemPath(fmt.Sprintf("%s/%s", configDir, keysFilename))
	if data, err := ioutil.ReadFile(keysPath); err == nil {
		keyData := keys{}

		if err = json.Unmarshal(data, &keyData); err == nil {
			Token = keyData.Token
		}
	}
}

// RequireConfig is a middleware that makes sure all of the config variables
// are initialized before we reach the action. We do this instead of just using
// a file func init() to prevent expensive operations when we don't need them.
func RequireConfig(action func(*cli.Context) error) func(*cli.Context) error {
	return func(c *cli.Context) error {
		getConfig()

		return action(c)
	}
}

// RequireHost requires the config and ensures that there is a remote host saved. If not,
// prompt for one before continuing
func RequireHost(action func(*cli.Context) error) func(*cli.Context) error {
	return func(c *cli.Context) error {
		getConfig()
		if Host == "" {
			iocli.Error("Your remote host is not set")

			var inp iocli.PromptInput
			isFirst := true
			host := ""

			for {
				if !isFirst {
					iocli.Error("Uh oh! Looks like the host %s is either", host)
					iocli.Error("unreachable or the server is not up and running")

					inp = iocli.PromptRune("Would you like to enter a different host? [Y/n]")
					if inp.IsNo() {
						os.Exit(0)
					}
				} else {
					isFirst = false
				}

				inp = iocli.PromptHost("Enter your remote host")
				host = inp.Response

				if checkHost(host) {
					break
				}
			}

			cfg := UserConfig{Host: host}
			Host = host
			SaveConfig(cfg)
		}

		return action(c)
	}
}

func checkHost(host string) bool {
	path := fmt.Sprintf("%s/api/core/status", host)
	client := http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(path)
	if err != nil {
		return false
	}

	return resp.StatusCode == http.StatusOK
}

// RequireAuth is a middleware to make sure you are logged in
func RequireAuth(action func(*cli.Context) error) func(*cli.Context) error {
	return func(c *cli.Context) error {
		getConfig()

		if Token == "" {
			iocli.Error("You must be logged in to do this action")
			iocli.Error("Run `gitup login` to log in or ")
			iocli.Error("run `gitup signup` to create an account")

			return nil
		}

		reqHost := RequireHost(func(cl *cli.Context) error {
			return action(cl)
		})

		return reqHost(c)
	}
}

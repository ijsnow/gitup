package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"

	"github.com/ijsnow/gitup/iocli"
)

// UserConfig is used for saving an updated config
type UserConfig struct {
	Uname string `json:"username"`
	Email string `json:"email"`
	Host  string `json:"host"`
}

// UserKeys is used for saving an updated config
type UserKeys struct {
	Token string `json:"token"`
}

// SaveConfig is used to save an updated config
func SaveConfig(uc UserConfig, override ...bool) error {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)

		return err
	}

	hdir := fromSystemPath(usr.HomeDir)

	configDir := fmt.Sprintf("%s/.gitup/", hdir)
	if _, err := os.Stat(toSystemPath(configDir)); os.IsNotExist(err) {
		os.Mkdir(toSystemPath(configDir), 0755)
	}

	getConfig()

	isOverride := len(override) > 0

	writeConfig := UserConfig{
		Uname: Uname,
		Email: Email,
	}

	if uc.Uname != "" || isOverride {
		writeConfig.Uname = uc.Uname
	}

	if uc.Email != "" || isOverride {
		writeConfig.Email = uc.Email
	}

	if uc.Host != "" {
		writeConfig.Host = uc.Host
	}

	if data, err := json.MarshalIndent(writeConfig, "", "  "); err == nil {
		configPath := fmt.Sprintf("%s/%s", configDir, configFilename)

		err = ioutil.WriteFile(toSystemPath(configPath), data, 0644)
		if err != nil {
			iocli.Error("Error saving config file: %s", err)

			return err
		}
	} else {
		iocli.Error("Error saving config file: %s", err)

		return err
	}

	return nil
}

// SaveKeys is used to save new keys
func SaveKeys(uk UserKeys, override ...bool) error {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)

		return err
	}

	hdir := fromSystemPath(usr.HomeDir)

	configDir := fmt.Sprintf("%s/.gitup/", hdir)
	if _, err := os.Stat(toSystemPath(configDir)); os.IsNotExist(err) {
		os.Mkdir(configDir, 0755)
	}

	getConfig()

	isOverride := len(override) > 0

	writeKeys := UserKeys{
		Token: Token,
	}

	if uk.Token != "" || isOverride {
		writeKeys.Token = uk.Token
	}

	if data, err := json.MarshalIndent(writeKeys, "", "  "); err == nil {
		keysPath := fmt.Sprintf("%s/%s", configDir, keysFilename)

		err = ioutil.WriteFile(toSystemPath(keysPath), data, 0644)
		if err != nil {
			iocli.Error("Error saving config file: %s", err)
			return err
		}
	} else {
		iocli.Error("Error saving config file: %s", err)
		return err
	}

	return nil
}

// Logout logs out by deleting the user config
func Logout() error {
	keys := UserKeys{
		Token: "",
	}

	err := SaveKeys(keys, true)
	if err != nil {
		return err
	}

	cnf := UserConfig{
		Uname: "",
		Email: "",
	}

	err = SaveConfig(cnf, true)
	if err != nil {
		return err
	}

	return nil
}

package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"

	"gitup.io/isaac/gitup/iocli"
)

// UserConfig is used for saving an updated config
type UserConfig struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Host     string `json:"host"`
}

// UserKeys is used for saving an updated config
type UserKeys struct {
	Token   string `json:"token"`
	Key     string `json:"key"`
	KeyPath string
}

// SaveConfig is used to save an updated config
func SaveConfig(uc UserConfig) error {
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

	writeConfig := UserConfig{
		Username: Username,
		Email:    Email,
	}

	if uc.Username != "" {
		writeConfig.Username = uc.Username
	}

	if uc.Email != "" {
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
func SaveKeys(uk UserKeys) error {
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

	writeKeys := UserKeys{
		Token: Token,
		Key:   Key,
	}

	if uk.Token != "" {
		writeKeys.Token = uk.Token
	}

	if uk.Key != "" {
		writeKeys.Key = uk.Key
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

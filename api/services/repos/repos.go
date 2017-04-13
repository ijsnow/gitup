package repos

import (
	"log"
	"os"
	"os/user"
	"regexp"

	"gitup.io/isaac/gitup/api/config"
)

var repoDir string

func init() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile("~")

	path := re.ReplaceAllLiteralString(config.App.RepoDir, usr.HomeDir)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}

	repoDir = path
}

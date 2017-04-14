package fspaths

import (
	"log"
	"os/user"
	"regexp"
)

// ExpandTildePath expands paths containing ~ to the full homedir
func ExpandTildePath(path string) string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile("^~")

	return re.ReplaceAllLiteralString(path, usr.HomeDir)
}

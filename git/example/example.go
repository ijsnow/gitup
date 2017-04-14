package main

import (
	"fmt"
	"os/user"

	git "github.com/ijsnow/gitup/git"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	path := fmt.Sprintf("%s/ws/tmp/git/test-1/test.git", usr.HomeDir)

	_, err = git.CreateBareRepo(path)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Repo created at %s\n", path)
}

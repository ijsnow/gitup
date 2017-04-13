package repos

import (
	"fmt"
	"os"
)

// TODO: Modify user permissions such as from the following from the node api
//
//	shell.exec(`echo 'asac413gu' | sudo -kS chgrp -R gitup ${userPath}`);
//	console.log('New user group access permissions modified.');

// CreateUserRepoDirectory creates the directories needed for repos when a user signs up
func CreateUserRepoDirectory(uname string) error {
	err := os.MkdirAll(fmt.Sprintf("%s/%s", repoDir, uname), os.ModePerm)
	if err != nil {
		return err
	}

	// Modify user perms here

	return nil
}

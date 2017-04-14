package auth

import (
	"net/http"

	"gitup.io/isaac/gitup/cli/api"
	"gitup.io/isaac/gitup/cli/config"
	"gitup.io/isaac/gitup/iocli"
	"gitup.io/isaac/gitup/types"
)

// Login logs the user in
func Login() error {
	u := &types.User{}
	requestSuccess := false

	for !requestSuccess {
		inp := iocli.PromptUname("Enter username")
		u.Uname = inp.Response

		inp = iocli.PromptPassword("Enter password")
		u.Password = inp.Response

		_, resp := api.Login(u)
		if resp.StatusCode == http.StatusNotFound {
			iocli.Error("Oops! Invalid Credentials")

			inp := iocli.PromptRune("Would you like to try again? [Y/n]")
			if inp.IsNo() {
				return nil
			}
		} else {
			requestSuccess = true
		}
	}

	keys := config.UserKeys{
		Token: u.Token,
	}

	err := config.SaveKeys(keys)
	if err != nil {
		iocli.Error("Error creating local session")
	}

	cnf := config.UserConfig{
		Uname: u.Uname,
		Email: u.Email,
	}

	err = config.SaveConfig(cnf)
	if err != nil {
		iocli.Error("Error creating local session")
	}

	iocli.Success("Successfully logged in as %s", u.Uname)

	return nil
}

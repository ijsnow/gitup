package auth

import (
	"net/http"

	"github.com/ijsnow/gitup/cli/api"
	"github.com/ijsnow/gitup/cli/config"
	"github.com/ijsnow/gitup/iocli"
	"github.com/ijsnow/gitup/types"
)

// Signup signs a user up
func Signup(u *types.User) error {
	var promptResp iocli.PromptInput

	requestSuccess := false

	for !requestSuccess {
		promptResp = iocli.PromptUname("What username would you like?")
		u.Uname = promptResp.Response

		promptResp = iocli.PromptEmail("What's your email?")
		u.Email = promptResp.Response

		promptResp = iocli.PromptPassword("Enter a password")
		u.Password = promptResp.Response

		iocli.Info("Signing up with:")
		iocli.Info("  username: %s", u.Uname)
		iocli.Info("  email:    %s", u.Email)

		ok, resp := api.Signup(u)
		if resp.StatusCode == http.StatusNotFound {
			iocli.Error("Oops! Invalid Credentials")

			inp := iocli.PromptRune("Would you like to try again? [Y/n]")
			if inp.IsNo() {
				return nil
			}
		} else if !ok {
			iocli.Error("Uh oh! There are a few issues with your new credentials:")
			iocli.Errors(resp.Errors)

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

	iocli.Success("Successfully signed up!")
	iocli.Success("You are now logged in as %s", u.Uname)

	inp := iocli.PromptRune("Would you like a quick tour? [Y/n]")
	if inp.IsYes() {
		iocli.Info("To create a new repo run `gitup repo new`")
		iocli.Info("If you have any more questions, run `gitup help`")
		iocli.Info("or visit the documenation site for more info")
	}

	iocli.Success("Happy coding!")

	return nil
}

package router

import (
	"net/http"

	"github.com/ijsnow/gittp"
	"github.com/ijsnow/gitup/api/controllers"
	"github.com/ijsnow/gitup/api/services/repos"
)

func getGitHandler() http.Handler {
	ctl := controllers.NewGitController()

	return gittp.NewHandler(repos.GetRepoDir(), ctl.ServeRepo)
}

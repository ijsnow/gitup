package router

import (
	"net/http"

	"github.com/ijsnow/gittp"
	"gitup.io/isaac/gitup/api/controllers"
	"gitup.io/isaac/gitup/api/services/repos"
)

func getGitHandler() http.Handler {
	ctl := controllers.NewGitController()

	return gittp.NewHandler(repos.GetRepoDir(), ctl.ServeRepo)
}

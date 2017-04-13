package router

import (
	"github.com/gorilla/mux"
	"gitup.io/isaac/gitup/api/controllers"
	httputils "gitup.io/isaac/gitup/utils/http"
)

func setGitRoutes(router *mux.Router) {
	ctl := controllers.NewGitController()
	sub := router.PathPrefix("/core/").Subrouter()

	sub.Handle("/", httputils.
		NewHandler(ctl.GetRepo)).
		Methods("GET")
}

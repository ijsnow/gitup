package router

import (
	"github.com/gorilla/mux"
	"gitup.io/isaac/gitup/api/controllers"
	"gitup.io/isaac/gitup/api/services/middlewares"
	httputils "gitup.io/isaac/gitup/utils/http"
)

func setRepoRoutes(router *mux.Router) {
	ctl := controllers.NewRepoController()
	sub := router.PathPrefix("/repo/").Subrouter()

	sub.Handle("/", httputils.NewHandler(ctl.CreateRepo).
		UseMiddlewares(middlewares.RequireAuth)).
		Methods("POST")
	sub.Handle("/", httputils.NewHandler(ctl.DeleteRepo).
		UseMiddlewares(middlewares.RequireAuth)).
		Methods("DELETE")
}

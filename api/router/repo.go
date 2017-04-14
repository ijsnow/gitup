package router

import (
	"github.com/gorilla/mux"
	"github.com/ijsnow/gitup/api/controllers"
	"github.com/ijsnow/gitup/api/services/middlewares"
	httputils "github.com/ijsnow/gitup/utils/http"
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

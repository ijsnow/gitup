package router

import (
	"github.com/gorilla/mux"
	"github.com/ijsnow/gitup/api/controllers"
	"github.com/ijsnow/gitup/api/services/middlewares"
	httputils "github.com/ijsnow/gitup/utils/http"
)

func setAuthRoutes(router *mux.Router) {
	ctl := controllers.NewAuthController()
	sub := router.PathPrefix("/auth/").Subrouter()

	sub.Handle("/login", httputils.NewHandler(ctl.Login)).
		Methods("POST")
	sub.Handle("/signup", httputils.NewHandler(ctl.CreateUser)).
		Methods("POST")
	sub.Handle("/logout", httputils.NewHandler(ctl.Logout).
		UseMiddlewares(middlewares.RequireAuth)).
		Methods("DELETE")
}

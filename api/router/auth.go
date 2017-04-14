package router

import (
	"github.com/gorilla/mux"
	"gitup.io/isaac/gitup/api/controllers"
	"gitup.io/isaac/gitup/api/services/middlewares"
	httputils "gitup.io/isaac/gitup/utils/http"
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

package router

import (
	"github.com/gorilla/mux"
	"gitup.io/isaac/gitup/api/controllers"
	httputils "gitup.io/isaac/gitup/utils/http"
)

func setAuthRoutes(router *mux.Router) {
	co := controllers.NewAuthController()
	sub := router.PathPrefix("/auth/").Subrouter()

	sub.Handle("/login", httputils.NewHandler(co.Login)).
		Methods("POST")
	sub.Handle("/signup", httputils.NewHandler(co.CreateUser)).
		Methods("POST")
}

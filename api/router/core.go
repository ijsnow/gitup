package router

import (
	"github.com/gorilla/mux"
	"github.com/ijsnow/gitup/api/controllers"
	httputils "github.com/ijsnow/gitup/utils/http"
)

func setCoreRoutes(router *mux.Router) {
	ctl := controllers.NewCoreController()
	sub := router.PathPrefix("/core/").Subrouter()

	sub.Handle("/status", httputils.
		NewHandler(ctl.Status)).
		Methods("GET")
}

package router

import (
	"github.com/gorilla/mux"
	"gitup.io/isaac/gitup/api/controllers"
	httputils "gitup.io/isaac/gitup/utils/http"
)

func setCoreRoutes(router *mux.Router) {
	ctl := controllers.NewCoreController()
	sub := router.PathPrefix("/core/").Subrouter()

	sub.Handle("/status", httputils.
		NewHandler(ctl.Status)).
		Methods("GET")
}

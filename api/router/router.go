package router

import (
	"io"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// InitRouter returns the router object
func InitRouter() http.Handler {
	r := mux.NewRouter()
	sub := r.PathPrefix("/api/").Subrouter()

	// Routes handling
	setCoreRoutes(sub)
	setAuthRoutes(sub)
	setGitRoutes(sub)

	sub.HandleFunc("/", rootHandler).Methods("GET")

	routerWithMiddlewares := handlers.LoggingHandler(os.Stdout, r)
	routerWithMiddlewares = handlers.RecoveryHandler()(routerWithMiddlewares)

	return routerWithMiddlewares
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, "gitup API\n")
}

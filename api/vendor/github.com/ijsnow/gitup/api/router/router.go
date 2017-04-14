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
	setRepoRoutes(sub)

	sub.HandleFunc("/", rootHandleFunc).Methods("GET")

	rh := rootHandler{
		git: getGitHandler(),
		web: getWebHandler(),
		api: r,
	}

	routerWithMiddlewares := handlers.LoggingHandler(os.Stdout, rh)
	routerWithMiddlewares = handlers.RecoveryHandler()(routerWithMiddlewares)

	return routerWithMiddlewares
}

func rootHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, "gitup API\n")
}

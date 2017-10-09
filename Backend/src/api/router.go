package api

import (
	"net/http"
	"api/decorators"
	"api/routes"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter()
	router.StrictSlash(true) // / make final slash optional in the uri

	allroutes := routes.Routes_main
	allroutes = append(allroutes, routes.TodoRoutes...)

	for _, route := range allroutes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = decorators.Cors(handler)
		handler = decorators.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}

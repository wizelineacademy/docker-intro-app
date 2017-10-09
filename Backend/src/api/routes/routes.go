package routes

import (
	"net/http"
	"api/handlers/simple"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

var Routes_main = Routes{
	Route{
		"Index",
		"GET",
		"/",
		simple.Index,
	},
	Route{
		"Healthz",
		"GET",
		"/healthz",
		simple.Healthz,
	},
}

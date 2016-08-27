package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/dimitrisCBR/shameboardAPI/v2/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
		Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handlers.Index,
	},
	Route{
		"Shames",
		"GET",
		"/shames",
		handlers.Shames,
	},
	Route{
		"Shame",
		"GET",
		"/shames/{shame_id}",
		handlers.Shame,
	},
	Route{
		"ShameCreate",
		"POST",
		"/shames/create",
		handlers.ShameCreate,
	},
}

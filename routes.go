package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route - route for restapi
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes - slice ot routes for rest api
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
		Index,
	},
	Route{
		"CellIndex",
		"GET",
		"/cells",
		CellIndex,
	},
	Route{
		"CellShow",
		"GET",
		"/cells/{CellID}",
		CellShow,
	},
}

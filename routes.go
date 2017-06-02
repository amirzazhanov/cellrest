package main

import (
	"net/http"
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
		"/cells/{CellType}",
		CellShow,
	},
	//	Route{
	//		"CellByType",
	//		"GET",
	//		"/cellstype/{CellType}",
	//		CellByType(session),
	//	},
}

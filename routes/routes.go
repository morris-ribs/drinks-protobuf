package routes

import (
	"net/http"

	"github.com/drinks-protobuf/handlers"
)

//Route defines a route from our server
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes defines an array of Route
type Routes []Route

//routes is the collection of routes used in our solution
var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handlers.Index,
	},
	Route{
		"SaveDrink",
		"POST",
		"/",
		handlers.SaveDrink,
	},
}

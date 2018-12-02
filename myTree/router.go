package myTree

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/grendach/myTree/logger"
)

var controller = &Controller{Repository: Repository{}}

// Create type for route
//  type Route define a route structure

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

// Routes define a list of routes for our API based on Route type
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		controller.Index,
	},
	Route{
		"AddPerson",
		"Post",
		"/",
		controller.AddPerson,
	},
	Route{
		"UpdatePerson",
		"PUT",
		"/",
		controller.UpdatePerson,
	},
	Route{
		"DeletePerson",
		"DELETE",
		"/{id}",
		controller.DeletePerson,
	},
}

// NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandleFunc
		handler = logger.Logger(handler, route.Name)
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}
	return router
}

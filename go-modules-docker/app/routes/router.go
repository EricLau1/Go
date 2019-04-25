package routes

import (
	"goapp/app/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Pattern string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

var routes = []Route{
	Route{"/", http.MethodGet, controllers.GetHello},
	Route{"/", http.MethodPost, controllers.PostHello},
	Route{"/", http.MethodPut, controllers.PutHello},
	Route{"/", http.MethodPatch, controllers.PatchHello},
	Route{"/", http.MethodDelete, controllers.DeleteHello},
}

func NewRouter() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		r.HandleFunc(route.Pattern, route.Handler).Methods(route.Method)
	}
	return r
}

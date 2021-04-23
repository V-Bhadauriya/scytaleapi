package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Route struct
type Route struct {
	URI          string
	Method       string
	Handler      func(w http.ResponseWriter, r *http.Request)
	AuthRequired bool
}

// Load the routes
func Load() []Route {
	routes := basicRoutes
	return routes
}


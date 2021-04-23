
package router

import (
"github.com/gorilla/mux"
"github.com/sandatasystem/scytaleapi/api/router/routes"
)

// New routes

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.Load(r)
}

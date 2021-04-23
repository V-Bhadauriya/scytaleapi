

package routes

import (
"github.com/V-Bhadauriya/GO-RESTApi/api/controllers"
"net/http"
)

var basicRoutes = []Route{
	Route{
		URI:     "/api/v1/ping",
		Method:  http.MethodGet,
		Handler: controllers.Ping,
	},

}

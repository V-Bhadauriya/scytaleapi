
package controllers

import (
	"net/http"
	_ "encoding/json"
	"github.com/gorilla/mux"
	"github.com/sandatasystem/scytaleapi/api/database"
	"github.com/sandatasystem/scytaleapi/api/models"
	"github.com/sandatasystem/scytaleapi/api/repository"
	"github.com/sandatasystem/scytaleapi/api/repository/crud"
	"github.com/sandatasystem/scytaleapi/api/responses"

)

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

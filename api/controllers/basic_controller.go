
package controllers

import (
	"net/http"
	_ "encoding/json"
	"github.com/gorilla/mux"
	"github.com/V-Bhadauriya/scytaleapi/api/database"
	"github.com/V-Bhadauriya/scytaleapi/api/models"
	"github.com/V-Bhadauriya/scytaleapi/api/repository"
	"github.com/V-Bhadauriya/scytaleapi/api/repository/crud"
	"github.com/V-Bhadauriya/scytaleapi/api/responses"

)

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

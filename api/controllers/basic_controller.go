
package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	 _ "encoding/json"
	 _ "github.com/gorilla/mux"
	 "github.com/V-Bhadauriya/scytaleapi/api/database"
	  "github.com/V-Bhadauriya/scytaleapi/api/models"
	 "github.com/V-Bhadauriya/scytaleapi/api/repository"
	 "github.com/V-Bhadauriya/scytaleapi/api/repository/crud"
	 "github.com/V-Bhadauriya/scytaleapi/api/responses"

)

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func GetWorkloads(w http.ResponseWriter, r *http.Request) {

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryWorkloadsCRUD(db)

	func(workloadRepository repository.WorkloadRepository) {
		workloads, err := workloadRepository.FindAll()
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, workloads)
	}(repo)
}


func CreateWorkload(w http.ResponseWriter, r *http.Request) {
	workload := models.Workload{}
	err := json.NewDecoder(r.Body).Decode(&workload)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	//err = workload.Validate()
	//if err != nil {

	//	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	//	return
	//}

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryWorkloadsCRUD(db)

	func(workloadRepository repository.WorkloadRepository) {

		workload, err := workloadRepository.Save(workload)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, workload.SpiffeID))
		responses.JSON(w, http.StatusCreated, workload)
	}(repo)
}

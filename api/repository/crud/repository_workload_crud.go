package crud

import (

"github.com/jinzhu/gorm"
"github.com/V-Bhadauriya/scytaleapi/api/models"
"github.com/V-Bhadauriya/scytaleapi/api/utils/channels"
)

// RepositoryPostsCRUD is the struct for the Post CRUD
type RepositoryWorkloadsCRUD struct {
	db *gorm.DB
}

// NewRepositoryPostsCRUD returns a new repository with DB connection
func NewRepositoryWorkloadsCRUD(db *gorm.DB) *RepositoryWorkloadsCRUD {
	return &RepositoryWorkloadsCRUD{db}
}

// Save returns a new post created or an error
func (r *RepositoryWorkloadsCRUD) Save(workload models.Workload) (models.Workload, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Workload{}).Create(&workload).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return workload, nil
	}
	return models.Workload{}, err
}

func (r *RepositoryWorkloadsCRUD) FindAll() ([]models.Workload, error) {
	var err error
	workloads := []models.Workload{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Find(&workloads).Error
		if err != nil {
			ch <- false
			return
		}

		ch <- true
	}(done)
	if channels.OK(done) {
		return workloads, nil
	}
	return nil, err
}


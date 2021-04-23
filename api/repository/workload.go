package repository

import (
"github.com/V-Bhadauriya/GO-RESTApi/api/models"
)

// PostRepository is the interface Post CRUD
type WorkloadRepository interface {
	Save(models.Workload) (models.Workload, error)
	// FindAll(uint32) ([]models.Job, error)
	// FindByID(uint32, uint32) (models.Job, error)
	// Update(uint32, uint32, models.Job) (int64, error)
	// Delete(uint32, uint64) (int64, error)
}

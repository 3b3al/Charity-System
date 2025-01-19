package profilingRepos

import (
	"github.com/3b3al/Charity-System/core/profiling"
	"gorm.io/gorm"
)

type PatientRepo struct {
	db *gorm.DB
}

func NewPatientRepo(db *gorm.DB) profiling.IPatientRepo {
	return &PatientRepo{db: db}
}

func (r *PatientRepo) CreatePatient(patient *profiling.Patient) (*profiling.Patient, error) {
	err := r.db.Create(patient).Error
	if err != nil {
		return nil, err
	}
	return patient, err
}

func (r *PatientRepo) DeletePatientByID(id string) error {
	return r.db.Delete(&profiling.Patient{}, id).Error
}

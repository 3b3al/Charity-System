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

func (patientRepo PatientRepo) GetPatientByID(id string) (*profiling.Patient, error) {
	var patient profiling.Patient
	err := patientRepo.db.Where("id = ?", id).First(&patient).Error
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

func (patientRepo PatientRepo) CreateOrUpdatePatient(patient *profiling.Patient) (*profiling.Patient, error) {
	err := patientRepo.db.Save(patient).Error
	return patient, err
}

func (patientRepo PatientRepo) DeletePatientByID(id string) error {
	return patientRepo.db.Where("id = ?", id).Delete(&profiling.Patient{}).Error
}

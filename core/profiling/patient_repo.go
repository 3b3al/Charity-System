package profiling

import (
	"gorm.io/gorm"
)

type PatientRepo struct {
	DB *gorm.DB
}

func (patientRepo PatientRepo) CreatePatient(patient *Patient) error {
	return patientRepo.DB.Create(patient).Error
}

func (patientRepo PatientRepo) GetPatientByID(id string) (*Patient, error) {
	var patient Patient
	err := patientRepo.DB.Where("id = ?", id).First(&patient).Error
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

func (patientRepo PatientRepo) UpdatePatient(patient *Patient) error {
	return patientRepo.DB.Save(patient).Error
}

func (patientRepo PatientRepo) DeletePatient(id string) error {
	return patientRepo.DB.Where("id = ?", id).Delete(&Patient{}).Error
}

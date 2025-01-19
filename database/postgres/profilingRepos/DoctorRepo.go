package profilingRepos

import (
	"github.com/3b3al/Charity-System/core/profiling"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type DoctorRepo struct {
	db *gorm.DB
}

func NewDoctorRepo(db *gorm.DB) profiling.IDoctorRepo {
	return &DoctorRepo{db: db}
}

func (doctorRepo DoctorRepo) GetDoctor(id uuid.UUID) (*profiling.Doctor, error) {
	var doctor profiling.Doctor
	err := doctorRepo.db.Where("id = ?", id).First(&doctor).Error
	if err != nil {
		return nil, err
	}
	return &doctor, nil
}

func (doctorRepo DoctorRepo) CreateOrUpdateDoctor(doctor *profiling.Doctor) (*profiling.Doctor, error) {
	err := doctorRepo.db.Save(doctor).Error
	return doctor, err
}

func (doctorRepo DoctorRepo) DeleteDoctor(id uuid.UUID) error {
	return doctorRepo.db.Where("id = ?", id).Delete(&profiling.Doctor{}).Error
}

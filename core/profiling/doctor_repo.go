package profiling

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type DoctorRepo struct {
	DB *gorm.DB
}

func (doctorRepo DoctorRepo) CreateDoctor(doctor *Doctor) error {
	return doctorRepo.DB.Create(doctor).Error
}

func (doctorRepo DoctorRepo) GetDoctor(id uuid.UUID) (*Doctor, error) {
	var doctor Doctor
	err := doctorRepo.DB.Where("id = ?", id).First(&doctor).Error
	if err != nil {
		return nil, err
	}
	return &doctor, nil
}

func (doctorRepo DoctorRepo) UpdateDoctor(doctor *Doctor) error {
	return doctorRepo.DB.Save(doctor).Error
}

func (doctorRepo DoctorRepo) DeleteDoctor(id uuid.UUID) error {
	return doctorRepo.DB.Where("id = ?", id).Delete(&Doctor{}).Error
}

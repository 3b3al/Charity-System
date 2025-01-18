package profiling

import "gorm.io/gorm"

type gardianRepo struct {
	DB *gorm.DB
}

func (gardianRepo gardianRepo) CreateGuardian(guardian *Guardian) error {
	return gardianRepo.DB.Create(guardian).Error
}

func (gardianRepo gardianRepo) GetGuardianByID(id string) (*Guardian, error) {
	var guardian Guardian
	err := gardianRepo.DB.Where("id = ?", id).First(&guardian).Error
	if err != nil {
		return nil, err
	}
	return &guardian, nil
}

func (gardianRepo gardianRepo) UpdateGuardian(guardian *Guardian) error {
	return gardianRepo.DB.Save(guardian).Error
}

func (gardianRepo gardianRepo) DeleteGuardian(id string) error {
	return gardianRepo.DB.Where("id = ?", id).Delete(&Guardian{}).Error
}

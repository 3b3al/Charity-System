package profilingRepos

import (
	"github.com/3b3al/Charity-System/core/profiling"
	"gorm.io/gorm"
)

type gardianRepo struct {
	db *gorm.DB
}

func NewGardianRepo(db *gorm.DB) profiling.IGuardianRepo {
	return &gardianRepo{db: db}
}

func (gardianRepo gardianRepo) GetGuardianByID(id string) (*profiling.Guardian, error) {
	var guardian profiling.Guardian
	err := gardianRepo.db.Where("id = ?", id).First(&guardian).Error
	if err != nil {
		return nil, err
	}
	return &guardian, nil
}

func (gardianRepo gardianRepo) CreateOrUpdateGuardian(guardian *profiling.Guardian) (*profiling.Guardian, error) {
	err := gardianRepo.db.Save(guardian).Error
	return guardian, err
}

func (gardianRepo gardianRepo) DeleteGuardian(id string) error {
	return gardianRepo.db.Where("id = ?", id).Delete(&profiling.Guardian{}).Error
}

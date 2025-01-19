package profilingRepos

import (
	"github.com/3b3al/Charity-System/core/profiling"
	"gorm.io/gorm"
)

type AddressRepo struct {
	db *gorm.DB
}

func NewAddressRepo(db *gorm.DB) profiling.IAddressRepo {
	return &AddressRepo{db: db}
}

func (addressRepo *AddressRepo) GetAddressByID(id string) (*profiling.Address, error) {
	var address profiling.Address
	err := addressRepo.db.Where("id = ?", id).First(&address).Error
	if err != nil {
		return nil, err
	}
	return &address, nil
}

func (addressRepo *AddressRepo) CreateOrUpdateAddress(address *profiling.Address) error {
	return addressRepo.db.Save(&address).Error
}

func (addressRepo *AddressRepo) DeleteAddress(id string) error {
	return addressRepo.db.Where("id = ?", id).Delete(&profiling.Address{}).Error
}

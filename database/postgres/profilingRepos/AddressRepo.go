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

func (a *AddressRepo) CreateAddress(address *profiling.Address) (*profiling.Address, error) {
	err := a.db.Create(address).Error
	if err != nil {
		return nil, err
	}
	return address, err
}

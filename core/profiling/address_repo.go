package profiling

import "gorm.io/gorm"

type addressRepo struct {
	DB *gorm.DB
}

func (addressRepo addressRepo) CreateAddress(address *Address) error {
	return addressRepo.DB.Create(address).Error
}

func (addressRepo addressRepo) GetAddressByID(id string) (*Address, error) {
	var address Address
	err := addressRepo.DB.Where("id = ?", id).First(&address).Error
	if err != nil {
		return nil, err
	}
	return &address, nil
}

func (addressRepo addressRepo) UpdateAddress(address *Address) error {
	return addressRepo.DB.Save(address).Error
}

func (addressRepo addressRepo) DeleteAddress(id string) error {
	return addressRepo.DB.Where("id = ?", id).Delete(&Address{}).Error
}

package profiling

import "gorm.io/gorm"

type contactRepo struct {
	DB *gorm.DB
}

func (contactRepo contactRepo) CreateContact(contact *Contact) error {
	return contactRepo.DB.Create(contact).Error
}

func (contactRepo contactRepo) GetContactByID(id string) (*Contact, error) {
	var contact Contact
	err := contactRepo.DB.Where("id = ?", id).First(&contact).Error
	if err != nil {
		return nil, err
	}
	return &contact, nil
}

func (contactRepo contactRepo) UpdateContact(contact *Contact) error {
	return contactRepo.DB.Save(contact).Error
}

func (contactRepo contactRepo) DeleteContact(id string) error {
	return contactRepo.DB.Where("id = ?", id).Delete(&Contact{}).Error
}

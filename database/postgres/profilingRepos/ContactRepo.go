package profilingRepos

import (
	"github.com/3b3al/Charity-System/core/profiling"
	"gorm.io/gorm"
)

type ContactRepo struct {
	db *gorm.DB
}

func NewContactRepo(db *gorm.DB) profiling.IContactRepo {
	return &ContactRepo{db: db}
}

func (contactRepo *ContactRepo) GetContactByID(id string) (*profiling.Contact, error) {
	var contact profiling.Contact
	err := contactRepo.db.Where("id = ?", id).First(&contact).Error
	if err != nil {
		return nil, err
	}
	return &contact, nil
}

func (contactRepo *ContactRepo) CreateOrUpdateContact(contact *profiling.Contact) (*profiling.Contact, error) {
	err := contactRepo.db.Save(contact).Error
	return contact, err
}

func (contactRepo *ContactRepo) DeleteContact(id string) error {
	return contactRepo.db.Where("id = ?", id).Delete(&profiling.Contact{}).Error
}

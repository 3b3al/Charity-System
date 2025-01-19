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

func (c ContactRepo) CreateContact(contact *profiling.Contact) (*profiling.Contact, error) {
	err := c.db.Create(contact).Error
	if err != nil {
		return nil, err
	}

	return contact, nil
}

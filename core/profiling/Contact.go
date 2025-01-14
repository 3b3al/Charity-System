package profiling

import (
	"database/sql"
	"time"
)

type RelatedType string
const (
	PatientType  RelatedType = "patient"
	GuardianType RelatedType = "guardian"
	DoctorType   RelatedType = "doctor"
)

type ContactType string
const (
	PrimaryContact   ContactType = "primary"
	SecondaryContact ContactType = "secondary"
	EmergencyContact ContactType = "emergency"
)

type Contact struct {
	ID                   string       `gorm:"column:id;primaryKey"`
	RelatedID            string       `gorm:"column:related_id" `
	RelatedType          RelatedType  `gorm:"column:related_type" `
	PhoneNumer           string       `gorm:"column:phone_number" `
	AlternatePhoneNumber string       `gorm:"column:alternate_phone_number"`
	Email                string       `gorm:"column:email"`
	ContactType          ContactType  `gorm:"column:contact_type"`
	CreatedAt            time.Time    `gorm:"column:created_at"`
	UpdateAt             sql.NullTime `gorm:"column:updated_at"`
	DeleteAt             sql.NullTime `gorm:"column:deleted_at"`
}

type IContactRepo interface {
	CreateContact(contact *Contact) error
	GetContactByID(id string) (*Contact, error)
}

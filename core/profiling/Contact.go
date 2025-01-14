package profiling

import "github.com/gofrs/uuid"

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
	BaseModel
	RelatedID            uuid.UUID   `gorm:"type:uuid;column:related_id"`
	RelatedType          RelatedType `gorm:"type:varchar(100);column:related_type"`
	PhoneNumber          string      `gorm:"type:varchar(250);column:phone_number"`
	AlternatePhoneNumber string      `gorm:"type:varchar(250);column:alternate_phone_number"`
	Email                string      `gorm:"type:varchar(250);column:email"`
	ContactType          ContactType `gorm:"type:varchar(100);column:contact_type"`
}

type IContactRepo interface {
	CreateContact(contact *Contact) error
	GetContactByID(id string) (*Contact, error)
}

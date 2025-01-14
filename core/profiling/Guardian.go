package profiling

import (
	"time"

	"github.com/gofrs/uuid"
)

type Relationship string

const (
	Father  Relationship = "father"
	Mother  Relationship = "mother"
	Brother Relationship = "brother"
	Sister  Relationship = "sister"
	Other   Relationship = "other"
)

type Guardian struct {
	BaseModel
	FirstName    string       `gorm:"type:varchar(100);column:first_name"`
	LastName     string       `gorm:"type:varchar(100);column:last_name"`
	BirthDate    time.Time    `gorm:"column:birth_date"`
	PatientID    uuid.UUID    `gorm:"type:uuid;column:patient_id"`
	Relationship Relationship `gorm:"type:varchar(100);column:relationship"`
	AddressID    uuid.UUID    `gorm:"type:uuid;column:address_id"`
	Address      Address      `gorm:"foreignKey:AddressID"`
	Contacts     []Contact    `gorm:"foreignKey:RelatedID;references:ID"`
}

type IGuardianRepo interface {
	CreateGuardian(guardian *Guardian) error
}

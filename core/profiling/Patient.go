package profiling

import (
	"time"

	"github.com/gofrs/uuid"
)

type Patient struct {
	BaseModel
	FirstName string     `gorm:"type:varchar(100);column:first_name"`
	LastName  string     `gorm:"type:varchar(100);column:last_name"`
	BirthDate time.Time  `gorm:"column:birth_date"`
	AddressID uuid.UUID  `gorm:"type:uuid;column:address_id"`
	Address   Address    `gorm:"foreignKey:AddressID"`
	Contacts  []Contact  `gorm:"foreignKey:RelatedID;references:ID"`
	Guardians []Guardian `gorm:"foreignKey:PatientID;references:ID"`
	Notes     string     `gorm:"column:notes"`
	DoctorID  uuid.UUID  `gorm:"type:uuid;column:doctor_id"`
	Doctor    Doctor     `gorm:"foreignKey:DoctorID"`
}

type IPatientRepo interface {
	CreatePatient(patient *Patient) error
}

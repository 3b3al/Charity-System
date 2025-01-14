package profiling

import (
	"database/sql"
	"time"
)

type Patient struct {
	ID        string       `gorm:"column:id;primaryKey"`
	FirstName string       `gorm:"column:first_name"`
	LastName  string       `gorm:"column:last_name"`
	BirthDate time.Time    `gorm:"column:birth_date"`
	Address   Address      `gorm:"foreignKey:address_id"`
	Contacts  []Contact    `gorm:"foreignKey:related_id"`
	Guardians []Guardian   `gorm:"foreignKey:patient_id"`
	Notes     string       `gorm:"column:notes"`
	Doctor    Doctor       `gorm:"foreignKey:doctor_id"`
	CreatedAt time.Time    `gorm:"column:created_at"`
	UpdateAt  sql.NullTime `gorm:"column:updated_at"`
	DeleteAt  sql.NullTime `gorm:"column:deleted_at"`
}

type IPatientRepo interface {
	CreatePatient(patient *Patient) error
}

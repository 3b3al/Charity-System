package profiling

import (
	"database/sql"
	"time"
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
	ID           string       `gorm:"column:id;primaryKey"`
	FirstName    string       `gorm:"column:first_name"`
	LastName     string       `gorm:"column:last_name"`
	BirthDate    time.Time    `gorm:"column:birth_date"`
	Patient      Patient      `gorm:"column:patient_id"`
	Relationship Relationship `gorm:"column:relationship"`
	Address      Address      `gorm:"column:foreignKey:address_id"`
	Contacts     []Contact    `gorm:"foreignKey:related_id"`
	CreatedAt    time.Time    `gorm:"column:created_at"`
	UpdateAt     sql.NullTime `gorm:"column:updated_at"`
	DeleteAt     sql.NullTime `gorm:"column:deleted_at"`
}

type IGuardianRepo interface {
	CreateGuardian(guardian *Guardian) error
}

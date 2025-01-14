package profiling

import (
	"database/sql"
	"time"
)

type Doctor struct {
	ID             string       `gorm:"column:id;primaryKey"`
	FirstName      string       `gorm:"column:first_name"`
	LastName       string       `gorm:"column:last_name"`
	Specialization string       `gorm:"column:specialization"`
	Contacts       []Contact    `gorm:"foreignKey:related_id"`
	CreatedAt      time.Time    `gorm:"column:created_at"`
	UpdatedAt      sql.NullTime `gorm:"column:updated_at"`
	DeleteAt       sql.NullTime `gorm:"column:deleted_at"`
}

type IDoctorRepo interface {
	CreateDoctor(doctor *Doctor) error
}

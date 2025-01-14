package profiling

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        string `gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

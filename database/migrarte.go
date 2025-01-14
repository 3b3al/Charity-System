package database

import (
	"github.com/3b3al/Charity-System/core/profiling"
	"gorm.io/gorm"
)

var models = []interface{}{
	&profiling.Guardian{},
	&profiling.Patient{},
	&profiling.Doctor{},
	&profiling.Contact{},
	&profiling.Address{},
}

func Migrate(db *gorm.DB) error {
	for _, model := range models {
		err := db.AutoMigrate(model)
		if err != nil {
			return err
		}
	}
	return nil
}

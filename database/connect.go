package database

import (
	"log"

	
	"github.com/3b3al/Charity-System/Config"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(setting.GetDBDriver(), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   setting.Conf.Database.Schema + ".",
			SingularTable: false,
		},
	})
	if err != nil {
		return nil, err
	}
	log.Println("connected to database successfully")
	return db, nil
}

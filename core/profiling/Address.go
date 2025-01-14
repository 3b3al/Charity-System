package profiling

import (
	"database/sql"
	"time"
)

type Address struct {
	ID           string       `gorm:"column:id;primaryKey"`
	BuildingName string       `gorm:"column:building_name"`
	StreetName   string       `gorm:"column:street_name"`
	AreaName     string       `gorm:"column:area_name"`
	City         string       `gorm:"column:city"`
	Emirate      string       `gorm:"column:emirate"`
	PoBox        string       `gorm:"column:po_box"`
	Landmark     string       `gorm:"column:landmark"`
	Country      string       `gorm:"column:country;default:'United Arab Emirates'"`
	MakaniNumber string       `gorm:"column:makani_number"`
	CreatedAt    time.Time    `gorm:"column:created_at"`
	UpdatedAt    sql.NullTime `gorm:"column:updated_at"`
	DeleteAt     sql.NullTime `gorm:"column:deleted_at"`
}

type IAddressRepo interface {
	CreateAddress(address *Address) error
}

package profiling

type Address struct {
	BaseModel
	BuildingName string `gorm:"type:varchar(150);column:building_name"`
	StreetName   string `gorm:"type:varchar(150);column:street_name"`
	AreaName     string `gorm:"type:varchar(100);column:area_name"`
	City         string `gorm:"type:varchar(100);column:city"`
	Emirate      string `gorm:"type:varchar(100);column:emirate"`
	PoBox        string `gorm:"type:varchar(100);column:po_box"`
	Landmark     string `gorm:"type:varchar(100);column:landmark"`
	Country      string `gorm:"type:varchar(100);column:country;default:'United Arab Emirates'"`
	MakaniNumber string `gorm:"type:varchar(250);column:makani_number"`
}

type IAddressRepo interface {
	CreateAddress(address *Address) error
	GetAddressByID(id string) (*Address, error)
	UpdateAddress(address *Address) error
	DeleteAddress(id string) error
}

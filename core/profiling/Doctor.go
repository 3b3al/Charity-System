package profiling

type Doctor struct {
    BaseModel
    FirstName      string    `gorm:"type:varchar(100);column:first_name"`
    LastName       string    `gorm:"type:varchar(100);column:last_name"`
    Specialization string    `gorm:"type:varchar(100);column:specialization"`
    Contacts       []Contact `gorm:"foreignKey:RelatedID;references:ID"`
}

type IDoctorRepo interface {
	CreateDoctor(doctor *Doctor) error
}

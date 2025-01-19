package profilingRepos_test

import (
	"time"

	"github.com/3b3al/Charity-System/core/profiling"
	"github.com/3b3al/Charity-System/database"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

func SetupTestDB() (*gorm.DB, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetBaseModelObj() profiling.BaseModel {
	return profiling.BaseModel{
		ID:        uuid.Must(uuid.NewV4()),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}

func GetAddressObj() profiling.Address {
	return profiling.Address{
		BaseModel:    GetBaseModelObj(),
		City:         "Cairo",
		StreetName:   "El-Maadi",
		Emirate:      "Any Emirate",
		PoBox:        "any po box",
		Landmark:     "any landmark",
		Country:      "Egypt",
		MakaniNumber: "any makani number",
	}
}

func GetPatientObj() *profiling.Patient {
	baseModel := GetBaseModelObj()
	address := GetAddressObj()
	contact := GetContactObj()
	contact.RelatedID = baseModel.ID
	contact.RelatedType = profiling.PatientType
	doctor := GetDoctorObj()
	guardian := GetGuardianObj()
	guardian.Relationship = profiling.Father
	guardian.PatientID = baseModel.ID
	return &profiling.Patient{
		BaseModel: baseModel,
		FirstName: "John",
		LastName:  "Doe",
		BirthDate: time.Date(1998, time.May, 24, 0, 0, 0, 0, time.UTC),
		AddressID: address.ID,
		Address:   address,
		Contacts: []profiling.Contact{
			contact,
		},
		Guardians: []profiling.Guardian{
			guardian,
		},
		DoctorID: doctor.ID,
		Doctor:   doctor,
		Notes:    "any notes",
	}
}

func GetContactObj() profiling.Contact {
	return profiling.Contact{
		BaseModel:            GetBaseModelObj(),
		PhoneNumber:          "01000000000",
		AlternatePhoneNumber: "01111111111",
		Email:                "any@any.com",
		ContactType:          profiling.PrimaryContact,
	}
}

func GetGuardianObj() profiling.Guardian {
	baseModel := GetBaseModelObj()
	address := GetAddressObj()
	contact := GetContactObj()
	contact.RelatedID = baseModel.ID
	contact.RelatedType = profiling.GuardianType
	return profiling.Guardian{
		BaseModel: baseModel,
		FirstName: "any First Name",
		LastName:  "any Last Name",
		BirthDate: time.Date(1998, time.May, 24, 0, 0, 0, 0, time.UTC),
		AddressID: address.ID,
		Address:   address,
		Contacts: []profiling.Contact{
			contact,
		},
	}
}

func GetDoctorObj() profiling.Doctor {
	baseModel := GetBaseModelObj()
	contact := GetContactObj()
	contact.RelatedID = baseModel.ID
	contact.RelatedType = profiling.DoctorType
	return profiling.Doctor{
		BaseModel:      baseModel,
		FirstName:      "any First name",
		LastName:       "any Last Name",
		Specialization: "any Specialization",
		Contacts: []profiling.Contact{
			contact,
		},
	}
}

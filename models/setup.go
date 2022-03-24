package models

import (
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("../test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	// if err := database.SetupJoinTable(&Doctor{}, "Hospitals", &DoctorHospitals{}); err != nil {
	// 	log.Fatalf("Cannot create related models: %s", err)
	// }
	// if err := database.SetupJoinTable(&Hospital{}, "Doctors", &DoctorHospitals{}); err != nil {
	// 	log.Fatalf("Cannot create related models: %s", err)
	// }

	database.AutoMigrate(&Doctor{}, &Hospital{}, &Patient{}, &Reports{}, &Appointment{})

	DB = database
}

package models

import (
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "scar:passloll@tcp(10.46.144.2:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
  	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

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

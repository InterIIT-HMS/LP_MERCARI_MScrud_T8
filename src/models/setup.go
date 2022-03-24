package models

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	fmt.Println(dsn)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	} else {
		fmt.Println("Database connection established")
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

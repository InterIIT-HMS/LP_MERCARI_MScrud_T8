package models

import (
	"time"

	"gorm.io/gorm"
)

type Doctor struct {
	gorm.Model

	Name        string
	Degree      string
	Profession  string
	Experience  uint
	PhoneNumber string
	Hospitals   []*Hospital `gorm:"many2many:doctor_hospital;"`
}

type DoctorHospitals struct {
	CreatedAt  time.Time
	DoctorId   uint
	HospitalId uint
	Department string
}

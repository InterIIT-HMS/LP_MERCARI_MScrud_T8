package models

import "time"

type Doctor struct {
	DoctorID    uint64 `gorm:"primaryKey;autoIncrement:true"`
	Name        string
	Degree      string
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

package models

import (
	"gorm.io/gorm"
)

type Hospital struct {
	gorm.Model

	HospitalID  uint `gorm:"primaryKey;autoIncrement:true"`
	Name        string
	Address     string
	PhoneNumber string
	Rating      uint
	Doctors     []*Doctor `gorm:"many2many:doctor_hospital;"`
}

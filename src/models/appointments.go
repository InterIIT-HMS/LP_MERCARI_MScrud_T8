package models

import (
	"time"

	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model

	DoctorId    uint
	PatientId   uint
	Agenda      string
	DateAndTime time.Time
}

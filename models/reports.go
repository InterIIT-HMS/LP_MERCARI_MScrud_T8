package models

import (
	"time"

	"gorm.io/gorm"
)

type Reports struct {
	gorm.Model

	DoctorId       uint
	PatientId      uint
	HospitalId     uint
	ReportFiles    string
	Date           time.Time
	CombinedPdfUrl string
}

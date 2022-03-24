package models

import (
	"time"

	"gorm.io/gorm"
)

type Reports struct {
	gorm.Model

	ReportFiles    string
	Date           time.Time
	CombinedPdfUrl string
	Doctors        *[]Doctor   `gorm:"foreignkey:DoctorID"`
	Hospital       *[]Hospital `gorm:"foreignkey:HospitalID"`
	Patient        *Patient    `gorm:"foreignkey:PatientID"`
}

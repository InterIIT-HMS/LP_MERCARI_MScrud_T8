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
	Doctors        *[]Doctor
	Hospital       *[]Hospital
	Patient        *Patient
}

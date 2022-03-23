package models

import (
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model

	Name   string `gorm:"not null"`
	NHID   string `gorm:"unique"`
	Gender string
	Age    uint
}

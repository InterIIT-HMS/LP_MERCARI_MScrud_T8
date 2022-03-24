package models

type Patient struct {
	PatientID uint   `gorm:"primaryKey;autoIncrement:true"`
	Name      string `gorm:"not null"`
	NHID      string `gorm:"unique"`
	Gender    string
	Age       uint
}

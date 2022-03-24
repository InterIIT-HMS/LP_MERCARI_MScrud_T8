package models

type Hospital struct {
	HospitalID  uint `gorm:"primaryKey;autoIncrement:true"`
	Name        string
	Address     string
	PhoneNumber string
	Rating      uint
	Doctors     []*Doctor `gorm:"many2many:doctor_hospital;"`
}

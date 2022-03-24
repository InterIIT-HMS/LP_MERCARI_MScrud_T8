package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"crud/models"

	"github.com/gin-gonic/gin"
)

type Doctor struct {
	Name        string `json:"name"`
	Degree      string `json:"degree"`
	Profession  string `json:"profession"`
	Experience  uint   `json:"experience"`
	PhoneNumber string `json:"phone_number"`
	HospitalId  uint   `json:"hospital_id"`
}

type DoctorInput struct {
	Name        string
	Degree      string
	Profession  string
	Experience  uint
	PhoneNumber string
	Hospital    []models.Hospital
}

// GET /doctors
// Find all doctors
func FindDoctors(c *gin.Context) {
	var doctors []models.Doctor
	models.DB.Preload("Hospitals").Find(&doctors)

	c.JSON(http.StatusOK, gin.H{"data": doctors})
}

func FindDoctorById(id uint) (*models.Doctor, error) {
	var doctor *models.Doctor
	if err := models.DB.Preload("Hospitals").Where("doctor_id = ?", id).First(&doctor).Error; err != nil {
		return doctor, err
	}

	return doctor, nil
}

// GET /doctors/:id
// Find a doctor
func FindDoctor(c *gin.Context) {
	// Get model if exist
	var doctor models.Doctor
	doctor.DoctorID, _ = strconv.ParseUint(c.Param("id"), 10, 64)

	if err := models.DB.Preload("Hospitals").Where(&doctor).First(&doctor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": doctor})
}

// POST /doctors
// Create new doctor
func CreateDoctor(c *gin.Context) {
	// Validate input
	var input Doctor
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hospital := make([]*models.Hospital, 1)

	var err error

	hospital[0], err = FindHospitalById(input.HospitalId)
	fmt.Println(hospital[0])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create doctor
	doctor := models.Doctor{Name: input.Name, Degree: input.Degree, Profession: input.Profession, Experience: input.Experience, PhoneNumber: input.PhoneNumber, Hospitals: hospital}
	models.DB.Create(&doctor)

	c.JSON(http.StatusOK, gin.H{"data": doctor})
}

// PATCH /doctors/:id
// Update a doctor
func UpdateDoctor(c *gin.Context) {
	// Get model if exist
	var doctor models.Doctor
	if err := models.DB.Preload("Hospitals").Where("doctor_id = ?", c.Param("id")).First(&doctor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// // Validate input
	var input Doctor
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hospital := make([]*models.Hospital, 1)

	var err error

	hospital[0], err = FindHospitalById(input.HospitalId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// doctor.Hospitals = hospital
	doctorUpdate := models.Doctor{Name: input.Name, Degree: input.Degree, Profession: input.Profession, Experience: input.Experience, PhoneNumber: input.PhoneNumber, Hospitals: hospital}

	models.DB.Model(models.Doctor{}).Where(&doctor).Updates(&doctorUpdate)

	c.JSON(http.StatusOK, gin.H{"data": doctorUpdate})
}

// DELETE /doctors/:id
// Delete a doctor
func DeleteDoctor(c *gin.Context) {
	// Get model if exist
	var doctor models.Doctor
	doctor.DoctorID, _ = strconv.ParseUint(c.Param("id"), 10, 64)
	// if err := models.DB.Where("doctor_id = ?", c.Param("id")).First(&doctor).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	// 	return
	// }

	models.DB.Delete(&doctor)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

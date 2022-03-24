package controllers

import (
	"fmt"
	"net/http"

	"crud/models"

	"github.com/gin-gonic/gin"
)

type HospitalInput struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	Rating      uint   `json:"rating"`
	DoctorIds   []uint `json:"doctor_id`
}

// GET /hospitals
// Find all hospitals
func FindHospitals(c *gin.Context) {
	var hospitals []models.Hospital
	models.DB.Find(&hospitals)

	c.JSON(http.StatusOK, gin.H{"data": hospitals})
}

func FindHospitalById(id uint) (*models.Hospital, error) {
	// Get model if exist
	var hospital *models.Hospital
	if err := models.DB.Where("hospital_id = ?", id).First(&hospital).Error; err != nil {
		return hospital, err
	}

	return hospital, nil

}

// GET /hospitals/:id
// Find a hospital
func FindHospital(c *gin.Context) {
	// Get model if exist
	var hospital models.Hospital
	if err := models.DB.Where("hospital_id = ?", c.Param("id")).First(&hospital).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": hospital})
}

// POST /hospitals
// Create new hospital
func CreateHospital(c *gin.Context) {
	// Validate input
	var input HospitalInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	doctorsArr := make([]*models.Doctor, len(input.DoctorIds))
	for idx, id := range input.DoctorIds {
		doctorsArr[idx], _ = FindDoctorById(id)
	}

	// Create hospital
	hospital := models.Hospital{Name: input.Name, PhoneNumber: input.PhoneNumber, Address: input.Address, Doctors: doctorsArr}
	models.DB.Create(&hospital)

	c.JSON(http.StatusOK, hospital)
}

// PATCH /hospitals/:id
// Update a hospital
func UpdateHospital(c *gin.Context) {
	// Get model if exist
	println("*******")
	var hospital models.Hospital
	if err := models.DB.Where("hospital_id = ?", c.Param("id")).First(&hospital).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input HospitalInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(input)

	models.DB.Model(hospital).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": hospital})
}

// DELETE /hospitals/:id
// Delete a hospital
func DeleteHospital(c *gin.Context) {
	// Get model if exist
	var hospital models.Hospital

	if err := models.DB.Where(map[string]interface{}{
		"id": c.Param("id"),
	}).First(&hospital).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	println(&hospital)

	models.DB.Delete(&hospital)

	c.JSON(http.StatusOK, gin.H{"success": true})
}

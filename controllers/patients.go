package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wryonik/microservices/appointment/models"
)

type PatientInput struct {
	Name   string `json:"name"`
	NHID   string `json:"nhid"`
	Gender string `json:"gender"`
	Age    uint   `json:"age"`
}

// GET /patients
// Find all patients
func FindPatients(c *gin.Context) {
	var patients []models.Patient
	models.DB.Find(&patients)

	c.JSON(http.StatusOK, gin.H{"data": patients})
}

// GET /patients/:id
// Find a patient
func FindPatient(c *gin.Context) {
	// Get model if exist
	var patient models.Patient
	if err := models.DB.Where("id = ?", c.Param("id")).First(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": patient})
}

// POST /patients
// Create new patient
func CreatePatient(c *gin.Context) {
	// Validate input
	var input PatientInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create patient
	patient := models.Patient{Name: input.Name, NHID: input.NHID, Gender: input.Gender, Age: input.Age}
	models.DB.Create(&patient)

	c.JSON(http.StatusOK, gin.H{"data": patient})
}

// PATCH /patients/:id
// Update a patient
func UpdatePatient(c *gin.Context) {
	// Get model if exist
	println("*******")
	var patient models.Patient
	if err := models.DB.Where("id = ?", c.Param("id")).First(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input PatientInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(input)

	models.DB.Model(patient).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": patient})
}

// DELETE /patients/:id
// Delete a patient
func DeletePatient(c *gin.Context) {
	// Get model if exist
	var patient models.Patient

	if err := models.DB.Where(map[string]interface{}{
		"id": c.Param("id"),
	}).First(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	println(&patient)

	models.DB.Delete(&patient)

	c.JSON(http.StatusOK, gin.H{"success": true})
}

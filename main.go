package main

import (
	"log"

	"github.com/wryonik/microservices/appointment/controllers"
	"github.com/wryonik/microservices/appointment/models"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://b3e8e3572fd444c184b0ece249f8bd07@o1176298.ingest.sentry.io/6273808",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Doctor Routes
	r.GET("/doctor", controllers.FindDoctors)
	r.GET("/doctor/:id", controllers.FindDoctor)
	r.POST("/doctor", controllers.CreateDoctor)
	r.PATCH("/doctor", controllers.UpdateDoctor)
	r.DELETE("/doctor", controllers.DeleteDoctor)

	// Patient Routes
	r.GET("/patients", controllers.FindPatients)
	r.GET("/patients/:id", controllers.FindPatient)
	r.POST("/patients", controllers.CreatePatient)
	r.PATCH("/patients/:id", controllers.UpdatePatient)
	r.DELETE("/patients/:id", controllers.DeletePatient)

	// Hostpital Routes
	r.GET("/hospitals", controllers.FindHospitals)
	r.POST("/hospitals", controllers.CreateHospital)
	r.PATCH("/hospitals", controllers.UpdateHospital)
	r.DELETE("/hospitals", controllers.DeleteHospital)

	// Run the server
	r.Run()
}

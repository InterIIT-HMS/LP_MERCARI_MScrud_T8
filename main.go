package main

import (
	"crud/controllers"
	"crud/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Role  string `json:"given_name"`
	Email string `json:"email"`
	Id    string `json:"nickname"`
}

func authMid(c *gin.Context) {

	url := "https://dev-rgmfg73e.us.auth0.com/userinfo"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", c.Request.Header["Authorization"][0])

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	response := Response{}
	json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response.Email)
	fmt.Println(response.Role)
	fmt.Println(response.Id)
	c.Params = []gin.Param{
		{
			Key:   "email",
			Value: response.Email,
		},
		{
			Key:   "role",
			Value: response.Role,
		},
		{
			Key:   "id",
			Value: response.Id,
		},
		
	}
}

func main() {

	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	secureGroup := r.Group("/secure/", authMid)

	// Doctor Routes
	secureGroup.GET("/doctors", controllers.FindDoctors)
	secureGroup.GET("/doctor/:id", controllers.FindDoctor)
	r.POST("/doctor", controllers.CreateDoctor)
	secureGroup.PATCH("/doctor/:id", controllers.UpdateDoctor)
	secureGroup.DELETE("/doctor/:id", controllers.DeleteDoctor)

	// Patient Routes
	secureGroup.GET("/patients", controllers.FindPatients)
	secureGroup.GET("/patient/:id", controllers.FindPatient)
	r.POST("/patient", controllers.CreatePatient)
	secureGroup.PATCH("/patient/:id", controllers.UpdatePatient)
	secureGroup.DELETE("/patient/:id", controllers.DeletePatient)

	// Hostpital Routes
	secureGroup.GET("/hospitals", controllers.FindHospitals)
	secureGroup.GET("/hospital/:id", controllers.FindHospital)
	r.POST("/hospital", controllers.CreateHospital)
	secureGroup.PATCH("/hospital/:id", controllers.UpdateHospital)
	secureGroup.DELETE("/hospital/:id", controllers.DeleteHospital)

	// Run the server
	r.Run()
}

package main

import "C"
import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/models"
	"log"
	"net/http"
)

func getAllPeople(c *gin.Context) {
	data := models.AllPeople()
	c.JSON(http.StatusOK, data)
}

func getPeopleByID(c *gin.Context) {
	id := c.Param("id")
	fromString, err := uuid.FromString(id)
	if err != nil {
		c.JSON(http.StatusNotFound, "Data Not Found")
		return
	}
	data := models.FindPersonByID(fromString)
	if data == nil {
		c.JSON(http.StatusNotFound, "Data Not Found")
		return
	}
	c.JSON(http.StatusOK, data)
}

func main() {
	fmt.Println("SP// Backend Developer Test - RESTful Service")
	fmt.Println()

	// TODO: Add RESTful web service here
	router := gin.Default()
	router.GET("/people", getAllPeople)
	router.GET("/people/:id", getPeopleByID)

	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("gin run error: %v", err)
		return
	}

}

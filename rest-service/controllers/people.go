package controllers

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/models"
	"net/http"
)

// GetPeople GET /people
func GetPeople(c *gin.Context) {
	data := make([]*models.Person, 0)
	phoneNumber, gotPhoneNumber := c.GetQuery("phone_number")
	firstName, gotFirstName := c.GetQuery("first_name")
	lastName, gotLastName := c.GetQuery("last_name")
	// the priority need more detail of requirement
	if gotPhoneNumber {
		data = models.FindPeopleByPhoneNumber(phoneNumber)
	} else if gotFirstName && gotLastName {
		data = models.FindPeopleByName(firstName, lastName)
	} else {
		data = models.AllPeople()
	}
	c.JSON(http.StatusOK, data)
}

// GetPeopleByID GET /people/:id
func GetPeopleByID(c *gin.Context) {
	id := c.Param("id")
	fromString, err := uuid.FromString(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "data not found"})
		return
	}
	data := models.FindPersonByID(fromString)
	if data == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "data not found"})
		return
	}
	c.JSON(http.StatusOK, data)
}

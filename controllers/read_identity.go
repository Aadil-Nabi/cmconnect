package controllers

import (
	"net/http"

	"github.com/Aadil-Nabi/cmconnect/configs"
	"github.com/Aadil-Nabi/cmconnect/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// securityPin struct to store the value from the request.
type EmployeeDetails struct {
	Email       string `json:"email" binding:"required"`
	SecurityPin string `json:"securitypin" binding:"required"`
}

var employeeBody EmployeeDetails

func ReadPostHandler(c *gin.Context) {

	// ShouldBindJSON to bind the JSON data from the request body to the EmployeeDetails struct:
	c.ShouldBindJSON(&employeeBody)

	// Declare a variable of Identity model stored in DB
	var identity models.Identity

	// Get DB object
	DB := configs.ConnectDB()

	// Below DB query is to get the security pin and store in EmployeeDetails struct object.
	DB.Where("email=?", employeeBody.Email).First(&identity)

	// Lookup the requested user in the DB and store in a variable
	// Compare the Security Pin hash with the one store in DB
	// If Secuity Pin matches, we fetch the user details otherwise return
	err := bcrypt.CompareHashAndPassword([]byte(identity.SecurityPin), []byte(employeeBody.SecurityPin))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to authenticate security Pin, invalid security pin was provided",
		})
		return
	} else {
		res := DB.Where("email=?", employeeBody.Email).First(&identity)
		if res.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "identity not found",
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"result": identity,
			})
		}
	}

}

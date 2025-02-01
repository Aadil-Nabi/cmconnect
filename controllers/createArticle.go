package controllers

import (
	"net/http"

	"github.com/Aadil-Nabi/cmconnect/configs"
	"github.com/Aadil-Nabi/cmconnect/models"
	"github.com/gin-gonic/gin"
)

type identity struct {
	IdentityNumber string
	Department     string
}

func CreatePostHandler(c *gin.Context) {
	var identityPayload identity

	c.Bind(&identityPayload)

	identity := models.Identity{
		IdentityNumber: identityPayload.IdentityNumber,
		Department:     identityPayload.Department,
	}

	DB := configs.ConnectDB()
	DB.Create(&identity)

	c.JSON(http.StatusOK, gin.H{
		"result": "post created",
	})

}

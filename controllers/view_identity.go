package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type viewPayload struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var view viewPayload

func ViewIdentity(c *gin.Context) {

	c.Bind(&view)

	fmt.Println(view)

	c.JSON(http.StatusOK, gin.H{
		"payload": view,
	})

}

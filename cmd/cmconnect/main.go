package main

import (
	"fmt"

	"github.com/Aadil-Nabi/cmconnect/api/encryption"
	"github.com/Aadil-Nabi/cmconnect/configs"
	"github.com/Aadil-Nabi/cmconnect/controllers"
	"github.com/gin-gonic/gin"
)

func init() {
	configs.MustLoadEnvs()
}

func main() {
	fmt.Println("***********************************************************************************************")
	fmt.Println("**************************_________Welcome to the CM API_________******************************")
	fmt.Println("***********************************************************************************************")
	fmt.Println("***********************************************************************************************")

	fmt.Println()

	// CM APIS
	encryption.Encrypting()

	// Routers for

	// Create a gin router
	router := gin.Default()

	router.POST("/create", controllers.CreatePostHandler)
	router.GET("/read", controllers.ReadPostHandler)

	// Run the Server
	router.Run()

}

package main

import (
	"github.com/Aadil-Nabi/cmconnect/configs"
	"github.com/Aadil-Nabi/cmconnect/models"
)

// init function is a unique function that load before the main function to load necessary configurations
func init() {
	configs.MustLoadEnvs()
}

// Used only once to migrate the model into an actual database on postgres
func main() {

	DB := configs.ConnectDB()
	DB.AutoMigrate(&models.Identity{})

}

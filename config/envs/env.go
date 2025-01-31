package envs

import (
	"log"

	"github.com/joho/godotenv"
)

// GetEnvs function is used to get the environment variables.
// this function used godotenv library to load and use the environment variable stored in .env file
// there are other libraries like viper or cobra that can be used, but for simplicity we use godotenv
func GetEnvs() map[string]string {

	// Read method will load and return the environment variables in a map
	myEnvs, err := godotenv.Read()
	if err != nil {
		log.Fatal(err)
	}

	return myEnvs

}

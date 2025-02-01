package configs

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB function is used to get or open the DB connection instance for any operations on the database.
func ConnectDB() *gorm.DB {
	dsn := os.Getenv("DSN")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("unable to connect with database:%v", err)
	}

	return DB
}

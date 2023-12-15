package initializers

import (
	"github.com/c1tad3l/backend-wedo/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDb() {
	var err error
	env, _ := config.LoadConfig()
	dsn := env.DatabaseUrl
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
}

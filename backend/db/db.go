package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/axelsomerseth/varsity-dev-challenge/backend/config"
)

var Connection *gorm.DB

func Connect() *gorm.DB {
	var err error

	if Connection != nil {
		return Connection
	}

	driver := sqlite.Open(config.ENV.DBName)

	if Connection, err = gorm.Open(driver, &gorm.Config{}); err != nil {
		log.Fatal(fmt.Errorf("failed to open database connection %w", err))
	}

	fmt.Println("Success: database connected.")

	return Connection
}

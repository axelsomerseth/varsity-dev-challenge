package backend

import (
	"fmt"
	"log"

	"github.com/axelsomerseth/varsity-dev-challenge/backend/config"
	"github.com/axelsomerseth/varsity-dev-challenge/backend/db"
	"github.com/axelsomerseth/varsity-dev-challenge/backend/router"
)

func Start() {
	// Load env variables
	config.Load()

	// Connect Database
	db.Connect()

	// Start router
	r := router.Setup()

	// Handle errors
	port := fmt.Sprintf(":%s", config.ENV.Port)
	if err := r.Run(port); err != nil {
		log.Fatal("failed to start the backend. %w", err)
	}

}

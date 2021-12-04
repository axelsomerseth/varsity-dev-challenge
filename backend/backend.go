package backend

import (
	"log"

	"github.com/axelsomerseth/varsity-dev-challenge/backend/config"
	"github.com/axelsomerseth/varsity-dev-challenge/backend/router"
)

func Start() {
	// Load env variables
	config.Load()

	// Connect Database

	// Start router
	r := router.Setup()

	// Handle errors
	if err := r.Run(":8080"); err != nil {
		log.Fatal("failed to start the backend. %w", err)
	}

}

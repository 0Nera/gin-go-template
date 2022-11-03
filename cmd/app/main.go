package main

import (
	"log"

	"github.com/0nera/gin-go-template/config"
	"github.com/0nera/gin-go-template/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}

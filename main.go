package main

import (
	"github.com/joho/godotenv"
	"github.com/nicholasanthonys/hexagonal-banking/app"
	"github.com/nicholasanthonys/hexagonal-banking/logger"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("Error loading .env file")
	}

	logger.Info("starting application")
	app.Start()
}

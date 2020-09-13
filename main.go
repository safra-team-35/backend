package main

import (
	"os"

	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/safra-team-35/backend/data"
	"github.com/safra-team-35/backend/server"

	"github.com/gin-contrib/cors"
	"github.com/safra-team-35/backend/service"
)

func main() {
	logger.Info("Reading the initial configs...")

	db, err := data.Connect()
	if err != nil {
		panic(err)
	}
	svc := service.New(db)
	server := server.InitServer(svc)
	logger.Info("About to start the application...")

	server.Use(cors.Default())

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	if err := server.Run(":" + port); err != nil {
		panic(err)
	}
}

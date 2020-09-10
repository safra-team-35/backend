package main

import (
	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/safra-team-35/backend/data"
	"github.com/safra-team-35/backend/server"

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

	if err := server.Run(":3000"); err != nil {
		panic(err)
	}
}

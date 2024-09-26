package main

import (
	"virus_mocker/app/internal/config"
	"virus_mocker/app/internal/routes"
	"virus_mocker/app/pkg/logger"

	"gorm.io/gorm"
)

type server struct {
	Logger *logger.Logger
	DB     *gorm.DB
	Config *config.Config
}

func main() {
	server := &server{
		Logger: logger.Init(),
	}

	if err := routes.New(); err != nil {
		panic(err)
	}

	server.Logger.Info("Server was started correctly")
}

package main

import (
	"fmt"
	"virus_mocker/app/internal/config"
	"virus_mocker/app/internal/routes"
	"virus_mocker/app/pkg/logger"
)

type server struct {
	Logger *logger.Logger
	Config *config.Config
}

func main() {

	sever := &server{
		Logger: logger.Init(),
	}
	config, err := config.Init()
	if err != nil {
		panic(err)
	}

	fmt.Println(config)

	if err := routes.New(); err != nil {
		panic(err)
	}

	sever.Logger.Info("Server was started correctly")
}

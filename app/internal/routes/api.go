package routes

import (
	"virus_mocker/app/internal/broker"
	"virus_mocker/app/internal/config"
	"virus_mocker/app/internal/db"
	"virus_mocker/app/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Api struct {
	db     db.Database
	config *config.Config
	broker *redis.Client
	logger *logger.Logger
}

func New() error {
	var err error

	redInit, err := broker.Init()
	if err != nil {
		return err
	}

	db, err := db.Init()
	if err != nil {
		return err
	}

	config, err := config.Init()
	if err != nil {
		panic(err)
	}

	api := &Api{
		db:     db,
		config: config,
		logger: logger.Init(),
		broker: redInit,
	}

	api.logger.Info("Redis initialized!")

	r := gin.Default()
	if err := api.router(r); err != nil {
		return err
	}
	if err := r.Run(":8080"); err != nil {
		return err
	}

	return nil
}

func (a *Api) router(r *gin.Engine) error {
	api := r.Group("/api_v1")
	{
		api.GET("/ping", a.Ping)
		kata := api.Group("/kata")
		{
			scanner := kata.Group("/scanner")
			{
				v1 := scanner.Group("/v1")
				{
					sensors := v1.Group("/sensors")
					{
						sensors.POST("/dd62a4ee-a00b-438c-b95a-58006b8f6056/scans", a.CreateFile)
					}
				}
			}
		}
	}

	return nil
}

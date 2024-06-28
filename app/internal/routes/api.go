package routes

import (
	"virus_mocker/app/internal/broker"
	"virus_mocker/app/internal/config"
	"virus_mocker/app/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Api struct {
	logger *logger.Logger
	config *config.Config
	broker *redis.Client
}

func New() error {
	var err error

	redInit, err := broker.Init()
	if err != nil {
		return err
	}
	api := &Api{
		logger: logger.Init(),
		broker: redInit,
	}

	api.logger.Info("Redis initialized!")

	api.config, err = config.Init()
	if err != nil {
		return err
	}
	r := gin.Default()
	if err := r.Run(":8080"); err != nil {
		return err
	}

	if err := api.router(r); err != nil {
		return err
	}

	return nil
}

func (a *Api) router(r *gin.Engine) error {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return nil
}

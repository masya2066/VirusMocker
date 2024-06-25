package routes

import (
	"virus_mocker/app/internal/config"
	"virus_mocker/app/pkg/logger"

	"github.com/gin-gonic/gin"
)

type api struct {
	Logger *logger.Logger
	Config *config.Config
}

func New() error {
	var err error
	api := &api{
		Logger: logger.Init(),
	}
	api.Config, err = config.Init()
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

func (a *api) router(r *gin.Engine) error {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return nil
}

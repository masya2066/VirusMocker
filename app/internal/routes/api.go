package routes

import (
	"time"
	"virus_mocker/app/internal/config"
	"virus_mocker/app/internal/db"
	"virus_mocker/app/pkg/logger"

	"github.com/gin-contrib/cors"
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

	// redInit, err := broker.Init()
	// if err != nil {
	// 	return err
	// }

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
		// broker: redInit,
	}

	api.logger.Info("Redis initialized!")

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Allow all origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
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
					// sensors := v1.Group("/sensors")
					{
						// instance := sensors.Group("/:sensor_id")
						{
							v1.POST("/scans", a.CreateFile)
							v1.GET("/scans", a.GetFiles)
							v1.DELETE("/scans/:scan_id", a.DeleteFile)
						}
					}
				}
			}
		}
	}

	return nil
}

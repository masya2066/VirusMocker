package db

import (
	"virus_mocker/app/internal/config"
	"virus_mocker/app/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dbConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

type server struct {
	Logger *logger.Logger
	Config *config.Config
}

var DB *gorm.DB

func initConfig() (*dbConfig, error) {
	cfg, err := config.Init()
	if err != nil {
		return nil, err
	}

	return &dbConfig{
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		Username: cfg.Database.Username,
		Password: cfg.Database.Password,
		Name:     cfg.Database.Name,
	}, nil
}

func Init() error {
	cfg, err := initConfig()
	if err != nil {
		return err
	}

	dsn := "host=" + cfg.Host + " user=" + cfg.Username + " password=" + cfg.Password + " dbname=" + cfg.Name + " port=" + string(cfg.Port) + " sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	log := logger.Init()

	if err := DB.AutoMigrate(); err != nil {
		return err
	}

	log.Info("Database was started correctly")
	return nil
}

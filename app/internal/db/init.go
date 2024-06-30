package db

import (
	"strconv"
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

type Database struct {
	*gorm.DB
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

func Init() (db Database, fail error) {
	cfg, err := initConfig()
	if err != nil {
		return Database{DB}, err
	}

	dsn := "host=" + cfg.Host + " user=" + cfg.Username + " password=" + cfg.Password + " dbname=" + cfg.Name + " port=" + strconv.Itoa(cfg.Port) + " sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	log := logger.Init()

	if err := migrate(); err != nil {
		return Database{DB}, err
	}

	log.Info("Database was started correctly")
	return Database{DB}, nil
}

func migrate() error {
	if err := DB.AutoMigrate(
		&KataFile{},
	); err != nil {
		return err
	}
	return nil
}

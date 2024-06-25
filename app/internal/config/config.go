package config

import (
	"fmt"
	"os"
	"virus_mocker/app/pkg/files"
	"virus_mocker/app/pkg/folders"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

func Init() (*Config, error) {
	if err := folders.CheckFolderExists("app/configs"); err != nil {
		if err := folders.Create("app/configs"); err != nil {
			return nil, err
		}
		defaultConfig := defaultConfig()
		data, err := yaml.Marshal(defaultConfig)
		if err != nil {
			return nil, fmt.Errorf("error marshalling default config: %w", err)
		}

		if err := files.Create("app/configs/config.yml"); err != nil {
			return nil, err
		}

		err = os.WriteFile("app/configs/config.yml", data, 0666)
		if err != nil {
			return nil, fmt.Errorf("error writing default config file: %w", err)
		}

		return defaultConfig, nil
	}

	data, err := os.ReadFile("app/configs/config.yml")
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, fmt.Errorf("error parsing config file: %w", err)
	}

	return &cfg, nil
}

func defaultConfig() *Config {
	return &Config{
		Server: struct {
			Port int `yaml:"port"`
		}{
			Port: 8080,
		},
		Database: struct {
			Host     string `yaml:"host"`
			Port     int    `yaml:"port"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
			Name     string `yaml:"name"`
		}{
			Host:     "localhost",
			Port:     5432,
			Username: "user",
			Password: "password",
			Name:     "mydb",
		},
	}
}

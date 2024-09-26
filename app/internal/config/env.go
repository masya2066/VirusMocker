package config

import "github.com/joho/godotenv"

func GetEnv(key string) error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	return nil
}

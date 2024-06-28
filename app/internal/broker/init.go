package broker

import (
	"virus_mocker/app/pkg/logger"

	"github.com/go-redis/redis/v8"
)

type broker struct {
	client *redis.Client
	logger *logger.Logger
}

func Init() (red *redis.Client, fail error) {
	rd := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := rd.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return rd, nil
}

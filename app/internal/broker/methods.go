package broker

import (
	"context"
)

var ctx = context.Background()

func (rd *broker) Push(key string, value string) error {
	err := rd.client.LPush(ctx, key, value).Err()
	if err != nil {
		return err
	}

	return nil
}

func (rd *broker) List(key string) (res []string, fail error) {
	result, err := rd.client.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		rd.logger.Error("Error with connect to redis!", err)
	}

	return result, nil
}

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

func (rd *broker) Range(key string) (res []string, fail error) {
	result, err := rd.client.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		rd.logger.Error("Error with connect to redis!", err)
	}

	return result, nil
}

func (rd *broker) MapRange(key string) map[string]string {
	res, err := rd.Range(key)
	if err != nil {
		rd.logger.Error("Error with connect to redis!", err)
	}

	resultMap := make(map[string]string)
	for _, value := range res {
		resultMap[value] = value
	}

	return resultMap
}

func (rd *broker) FindInHashMap(key string, searchValue string) error {
	resultMap := rd.MapRange(key)
	if resultMap == nil {
		rd.logger.Error("Error generating map from range")
	}

	for _, value := range resultMap {
		if value == searchValue {
			return nil
		}
	}

	return nil
}

func (rd *broker) AddKey(key string, value string) error {
	if _, err := rd.client.Append(ctx, key, value).Result(); err != nil {
		return err
	}

	return nil
}

func (rd *broker) GetKey(key string) (string, error) {
	res, err := rd.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return res, nil
}

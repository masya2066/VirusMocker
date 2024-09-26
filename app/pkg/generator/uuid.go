package generator

import "github.com/google/uuid"

func Uuid() (string, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	return uuid.String(), nil
}

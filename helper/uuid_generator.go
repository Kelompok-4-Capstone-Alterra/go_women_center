package helper

import (
	"github.com/google/uuid"
)

type UuidGenerator interface {
	GenerateUUID() (string, error)
}

type googleUUID struct {}

func NewGoogleUUID() googleUUID {
	return googleUUID{}
}

func (g googleUUID) GenerateUUID() (string, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}
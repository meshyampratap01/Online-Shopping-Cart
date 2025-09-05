package utils

import "github.com/google/uuid"

func NewUUID() (string,error) {
	id:=uuid.New().String()

	return id,nil
}

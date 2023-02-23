package service

import (
	"github.com/prattnj/fms-go/model"
	"math/rand"
)

// Generic 500 response, used in all services
var serverError = model.GenericResponse{Success: false, Message: "Internal server error"}

var defaultIDLength = 32

func generateID(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

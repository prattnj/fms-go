package service

import (
	"github.com/prattnj/fms-go/model"
	"math/rand"
	"time"
)

// Generic 500 response, used in all services
var serverErrorStr = "Internal server error"
var tokenErrorStr = "Error: bad token"
var r = rand.New(rand.NewSource(time.Now().UnixNano()))
var serverError = model.GenericResponse{Success: false, Message: serverErrorStr}

var defaultIDLength = 32

func generateID(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

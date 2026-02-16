package utils

import (
	"math/rand"
)

func RandomString(length int, chars string) string {
	result := make([]byte, length)
	for i := range length {
		randomIndex := rand.Intn(len(chars))
		result[i] = chars[randomIndex]
	}
	return string(result)
}

func RandomOtp(length int) string {
	return RandomString(length, "1234567890")
}

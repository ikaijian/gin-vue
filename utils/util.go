package utils

import (
	"math/rand"
	"time"
)

func RandomString(number int) string {

	var letters = []byte("abcdeASDFASDFfghiDSFDASEREGFGFjklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]byte, number)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

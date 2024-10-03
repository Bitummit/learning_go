package utils

import (
	"math/rand"
	"time"
)


func NewRandomString(size int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	rand_string := make([]rune, size)

	for i := range size {
		rand_string[i] = letters[rnd.Intn(len(letters))]
	}

	return string(rand_string)

}
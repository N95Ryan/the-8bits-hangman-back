package utils

import (
	"math/rand"
	"time"
)

func GenerateID() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, 16)
	for i := range b {
		b[i] = charset[r.Intn(len(charset))]
	}
	return string(b)
}

func RandomWord() string {
	words := []string{"PROGRAMMATION", "ORDINATEUR", "ALGORITHME", "DEVELOPPEUR", "LOGICIEL"}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return words[r.Intn(len(words))]
}

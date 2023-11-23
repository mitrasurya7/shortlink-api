package helpers

import (
	"math/rand"
	"time"
)

func GenerateRandomString(length int) string {
	// Karakter yang dapat digunakan dalam string acak
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Inisialisasi generator angka acak dengan seed berdasarkan waktu
	rand.Seed(time.Now().UnixNano())

	// Membuat string acak
	randomString := make([]byte, length)
	for i := range randomString {
		randomString[i] = charset[rand.Intn(len(charset))]
	}

	return string(randomString)
}

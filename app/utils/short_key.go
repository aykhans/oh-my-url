package utils

import (
	"strings"
)

const keyCharacters string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateKey(n int) string {
	var result strings.Builder
	base := len(keyCharacters)

	for n > 0 {
		n--
		result.WriteByte(keyCharacters[n%base])
		n /= base
	}

	key := result.String()
	runes := []rune(key)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

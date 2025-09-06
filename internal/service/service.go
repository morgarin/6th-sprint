package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func IsMorseInput(input string) bool {
	input = strings.TrimSpace(input)
	// ПРИВЕТ .-.
	for _, char := range input {
		if (string(char) == ".") || (string(char) == "-") {
			return true
		}
	}
	return false
}

func ReverseMorse(input string) string {
	if IsMorseInput(input) {
		return morse.ToText(input)
	}
	return morse.ToMorse(input)
}

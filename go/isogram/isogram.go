package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	letters := make(map[string]bool)

	for _, char := range word {
		if unicode.IsLetter(char) {
			c := strings.ToLower(string(char))

			if _, exists := letters[c]; exists {
				return false
			}

			letters[c] = true
		}
	}

	return true
}

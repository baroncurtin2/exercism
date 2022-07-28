package logs

import (
	"fmt"
	"unicode/utf8"
)

var appCodes = map[string]string{
	"U+2757": "recommendation",
	"U+1F50D": "search",
	"U+2600": "weather",
}

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, char := range log {
		strChar := fmt.Sprintf("%U", char)

		if app, ok := appCodes[strChar]; ok {
			return app
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	runes := []rune(log)

	for i, char := range runes {
		if char == oldRune {
			runes[i] = newRune
		}
	}

	return string(runes)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	runeNumber := utf8.RuneCountInString(log)

	return runeNumber <= limit
}

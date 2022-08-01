package scrabble

import (
	"fmt"
	"strings"
)

type Letters []string

type LetterValue struct {
	value   int
	letters Letters
}

var LetterValues = []LetterValue{
	{
		value:   1,
		letters: Letters{"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"},
	},
	{
		value:   2,
		letters: Letters{"D", "G"},
	},
	{
		value:   3,
		letters: Letters{"B", "C", "M", "P"},
	},
	{
		value:   4,
		letters: Letters{"F", "H", "V", "W", "Y"},
	},
	{
		value:   5,
		letters: Letters{"K"},
	},
	{
		value:   8,
		letters: Letters{"J", "X"},
	},
	{
		value:   10,
		letters: Letters{"Q", "Z"},
	},
}

func convertLetterValuesToMap(lvs []LetterValue) map[string]int {
	lMap := make(map[string]int)

	for _, lv := range lvs {
		for _, l := range lv.letters {
			lMap[l] = lv.value
		}
	}

	return lMap
}

func Score(word string) int {
	score := 0

	if len(word) <= 0 {
		return 0
	}

	mapLetters := convertLetterValuesToMap(LetterValues)

	for _, char := range word {
		score += mapLetters[strings.ToUpper(string(char))]
	}

	fmt.Println(word, score)
	return score
}

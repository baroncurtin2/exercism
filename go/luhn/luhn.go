package luhn

import (
	"fmt"
	"strconv"
	"strings"
)

type LuhnDigits []int

func (digits LuhnDigits) Sum() int {
	sum := 0

	for _, digit := range digits {
		sum += digit
	}

	return sum
}

func Valid(id string) bool {
	digits, err := GatherDigits(id)

	if err != nil {
		return false
	}

	if len(digits) <= 1 {
		return false
	}

	luhn := digits.Sum()

	if luhn%10 == 0 {
		return true
	}

	return false
}

func GatherDigits(id string) (LuhnDigits, error) {
	var digits []int
	stripped := strings.ReplaceAll(id, " ", "")
	counter := 1

	for i := len(stripped) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(stripped[i]))

		if err != nil {
			return []int{}, err
		}

		if counter%2 == 0 {
			digits = append(digits, LuhnDouble(digit))
		} else {
			digits = append(digits, digit)
		}

		counter++
	}

	fmt.Println(id, digits)
	return digits, nil
}

func LuhnDouble(digit int) int {
	if digit*2 > 9 {
		return digit*2 - 9
	}

	return digit * 2
}

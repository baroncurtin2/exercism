package raindrops

import "fmt"

// Convert Converts a number into a string that contains raindrop sounds
// corresponding to the factor of a number
// 3 => 'Pling' ; 5 => 'Plang', 7 => 'Plong'
func Convert(number int) string {
	raindrop := ""
	factors := []int{3, 5, 7}
	sounds := []string{"Pling", "Plang", "Plong"}

	for i, f := range factors {
		if number%f == 0 {
			raindrop += sounds[i]
		}
	}

	if raindrop == "" {
		raindrop = fmt.Sprintf("%d", number)
	}

	return raindrop
}

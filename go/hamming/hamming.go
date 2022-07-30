package hamming

import "errors"

// Distance calculates the 'Hamming Distance' between two DNA sequences
func Distance(a, b string) (int, error) {
	// if lengths are different, return error
	if len(a) != len(b) {
		return -1, errors.New("DNA sequences must be the same length to compare")
	}

	count := 0
	for i := range a {
		if a[i] != b[i] {
			count++
		}
	}

	return count, nil
}

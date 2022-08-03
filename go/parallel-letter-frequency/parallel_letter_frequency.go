package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	ch := make(chan FreqMap, len(l))
	result := FreqMap{}

	for _, ll := range l {
		go func(line string) {
			ch <- Frequency(line)
		}(ll)
	}

	for i := 0; i < len(l); i++ {
		fMap := <-ch

		for k, v := range fMap {
			result[k] += v
		}
	}

	return result
}

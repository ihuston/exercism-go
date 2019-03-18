// Package letter provides concurrent letter counting tools.
package letter

// ConcurrentFrequency counts letters from text in a concurrent way.
func ConcurrentFrequency(texts []string) FreqMap {
	c := make(chan FreqMap, 10)
	for _, t := range texts {
		go frequencyHandler(t, c)
	}

	return accumulate(texts, c)
}

// frequencyHandler pushes a Frequency result to a channel.
func frequencyHandler(text string, c chan FreqMap) {
	c <- Frequency(text)
}

// accumulate combines multiple FreqMaps from a channel.
func accumulate(texts []string, c chan FreqMap) FreqMap {
	results := FreqMap{}
	for range texts {
		m := <-c
		for k, v := range m {
			results[k] += v
		}
	}
	return results
}

package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type wordFrequency struct {
	Word      string
	Frequency int
}

func Top10(text string) []string {
	words := strings.Fields(text)

	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	frequencies := make([]wordFrequency, 0, len(wordCount))
	for word, count := range wordCount {
		frequencies = append(frequencies, wordFrequency{Word: word, Frequency: count})
	}

	sort.Slice(frequencies, func(i, j int) bool {
		if frequencies[i].Frequency != frequencies[j].Frequency {
			return frequencies[i].Frequency > frequencies[j].Frequency
		}
		return frequencies[i].Word < frequencies[j].Word
	})

	topTenWords := make([]string, 0, 10)
	for i := 0; i < len(frequencies) && i < 10; i++ {
		topTenWords = append(topTenWords, frequencies[i].Word)
	}

	return topTenWords
}

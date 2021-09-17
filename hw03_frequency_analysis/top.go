package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type wordFrequency struct {
	word  string
	count int
}

func getTextFrequency(text string) (result []wordFrequency) {
	frequencyMap := make(map[string]int)
	for _, word := range strings.Fields(text) {
		frequencyMap[word] += 1
	}
	for key, value := range frequencyMap {
		result = append(result, wordFrequency{
			word:  key,
			count: value,
		})
	}
	return
}

func Top10(text string) (result []string) {
	frequency := getTextFrequency(text)
	sort.Slice(frequency, func(i, j int) bool {
		if frequency[i].count == frequency[j].count {
			return frequency[i].word < frequency[j].word
		}
		return frequency[i].count > frequency[j].count
	})

	if len(frequency) > 10 {
		frequency = frequency[:10]
	}

	for _, value := range frequency {
		result = append(result, value.word)
	}
	return
}

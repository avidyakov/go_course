package hw03frequencyanalysis

import (
	"errors"
	"sort"
	"strings"
)

type wordFrequency struct {
	word  string
	count int
}

var emptyWord = errors.New("empty word")

func processWord(rawWord string) (string, error) {
	rawWord = strings.Trim(rawWord, ",.-")
	rawWord = strings.ToLower(rawWord)

	if len(rawWord) == 0 {
		return "", emptyWord
	}
	return rawWord, nil
}

func getTextFrequency(text string) (result []wordFrequency) {
	frequencyMap := make(map[string]int)
	for _, word := range strings.Fields(text) {
		processedWord, err := processWord(word)
		if err == nil {
			frequencyMap[processedWord] += 1
		}
	}
	for key, value := range frequencyMap {
		result = append(result, wordFrequency{
			word:  key,
			count: value,
		})
	}
	return
}

func sortFrequency(frequency []wordFrequency) {
	sort.Slice(frequency, func(i, j int) bool {
		if frequency[i].count == frequency[j].count {
			return frequency[i].word < frequency[j].word
		}
		return frequency[i].count > frequency[j].count
	})
}

func Top10(text string) (result []string) {
	frequency := getTextFrequency(text)
	sortFrequency(frequency)

	if len(frequency) > 10 {
		frequency = frequency[:10]
	}

	for _, value := range frequency {
		result = append(result, value.word)
	}
	return
}

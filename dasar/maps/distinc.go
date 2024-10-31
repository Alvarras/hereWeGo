package dasar

import "strings"

func countDistinctWords(message []string) int {
	wordMap := make(map[string]bool)
	for _, msg := range message {
		words := strings.Fields(strings.ToLower(msg))
		for _, word := range words {
			wordMap[word] = true
		}
	}
	return len(wordMap)
}

package dasar

func getNameCounts(names []string) map[rune]map[string]int {
	counts := make(map[rune]map[string]int)
	for _, name := range names {
		runes := []rune(name)
		if len(runes) == 0 {
			continue
		}

		firstRune := runes[0]

		if counts[firstRune] == nil {
			counts[firstRune] = make(map[string]int)
		}
		counts[firstRune][name]++

	}
	return counts
}

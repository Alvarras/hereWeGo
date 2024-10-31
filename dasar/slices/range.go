package dasar

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, m := range msg {
		for _, b := range badWords {
			if m == b {
				return i
			}
		}
	}
	return -1
}

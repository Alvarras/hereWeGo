package dasar

func getMessageCosts(messages []string) []float64 {
	costs := make([]float64, len(messages))
	for i, message := range messages {
		costs[i] = float64(len(message)) * .01
	}
	return costs
}

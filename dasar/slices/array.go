package dasar

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	message := [3]string{primary, secondary, tertiary}
	cost := [3]int{0, 0, 0}

	for i := 0; i < len(message); i++ {
		if i == 0 {
			cost[i] = len(message[i])
		} else {
			cost[i] = cost[i-1] + len(message[i])
		}
	}
	return message, cost
}

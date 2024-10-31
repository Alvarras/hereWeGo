package dasar

func maxMessages(thresh int) int {
	totalCost := 0
	message := 0
	for i := 0; totalCost+100+i <= thresh; i++ {
		totalCost += 100 + i
		message++
	}
	return message
}

package dasar

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	var result []float64
	maxDay := 0
	for _, c := range costs {
		if c.day > maxDay {
			maxDay = c.day
		}
	}

	result = append(result, make([]float64, maxDay+1)...)
	for _, c := range costs {
		result[c.day] += c.value
	}
	return result
}

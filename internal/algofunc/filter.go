package algofunc

func FilterUp(numbers []float64) []float64 {
	var ret []float64
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > numbers[i-1] {
			ret = append(ret, numbers[i])
		}
	}

	return ret
}

func FilterDown(numbers []float64) []float64 {
	var ret []float64
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < numbers[i-1] {
			ret = append(ret, numbers[i])
		}
	}

	return ret
}

package algofunc

func CalcEma(numbers []float64, n int) []float64 {
	m := len(numbers)
	α := float64(2) / float64(n+1)
	var ema []float64

	ema = append(ema, numbers[0])

	for i := 1; i < m; i++ {
		ema = append(
			ema,
			(α*numbers[i])+((1-α)*ema[i-1]),
		)
	}

	return ema
}

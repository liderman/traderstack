package algofunc

import "github.com/shopspring/decimal"

func CalcRsi(data []decimal.Decimal) decimal.Decimal {
	totalGain := decimal.New(0, 0)
	totalLoss := decimal.New(0, 0)

	for i := 1; i < len(data); i++ {
		previousClose := data[i]
		currentClose := data[i-1]

		difference := currentClose.Sub(previousClose)

		// difference >= 0
		if difference.GreaterThanOrEqual(decimal.Zero) {
			totalGain = totalGain.Add(difference)
		} else {
			totalLoss = totalLoss.Sub(difference)
		}
	}

	if totalLoss.IsZero() {
		return decimal.NewFromInt(100)
	}

	if totalGain.IsZero() {
		return decimal.NewFromInt(0)
	}

	rs := totalGain.Div(totalLoss.Abs())

	//  100 - 100/(1+rs)
	return decimal.NewFromInt(100).Sub(
		decimal.NewFromInt(100).Div(
			decimal.NewFromInt(1).Add(rs),
		),
	)
}

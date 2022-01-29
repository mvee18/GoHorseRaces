package economy

type Money float64

const takePercentage = 0.15

func CalculateOdds(bet Money, total Money) Money {
	total -= total * takePercentage
	leftOver := total - bet

	return leftOver / bet
}

package economy

import "math/rand"

type Money float64

type OddsPool struct {
	Wager     Money
	Frequency float64
}

const takePercentage = 0.15

func CalculateOdds(bet Money, total Money) Money {
	total -= total * takePercentage
	leftOver := total - bet

	return leftOver / bet
}

func DistributeMoney(o []OddsPool) {
	// Somewhere between 0-100000, but mostly around 50,000.
	TotalWager := rand.NormFloat64() * 100000

	
}

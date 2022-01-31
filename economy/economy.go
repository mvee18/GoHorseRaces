package economy

import (
	"horses/models"
	"math/rand"
	"sort"
)

type Money float64

type OddsPool struct {
	Index     int
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

func fractionalSpeed(total Money, hs []models.Horse) {
	oddsPools := make([]OddsPool, len(hs))
	totalSpeed := 0.0

	// Sum up speed.
	for _, v := range hs {
		totalSpeed += v.Speed
	}

	for i, v := range hs {
		op := OddsPool{
			Index:     i,
			Wager:     0.0,
			Frequency: v.Speed / totalSpeed,
		}

		oddsPools = append(oddsPools, op)
	}
}

func rouletteDistribution(total Money, o *[]OddsPool) {
	for i, v := range OddsPool {

	}
}

func NormalizeSpeed(h []models.Horse) {
	// We need to sort the horses by speed to find the maximum value.
	sort.SliceStable(h, func(i, j int) bool {
		return h[i].Speed < h[j].Speed
	})

	minimumSpeed := h[0].Speed

	for i := 0; i < len(h); i++ {
		h[i].Speed = h[i].Speed / minimumSpeed
	}
}

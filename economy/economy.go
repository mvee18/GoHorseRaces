package economy

import (
	"horses/models"
	"sort"
)

type Money float64

type OddsPool struct {
	Index     int
	Wager     Money
	Frequency float64
}

const takePercentage = 0.15

func CalculateOdds(bet Money, total Money) float64 {
	total -= total * takePercentage
	leftOver := total - bet

	return float64(leftOver) / float64(bet)
}

func DistributeMoney(r *models.Race) {
	// Somewhere between 0-100000, but mostly around 50,000.
	totalWager := models.R1.NormFloat64()*25000 + 50000

	fractionalSpeed(Money(totalWager), &r.Horses, CalculateOdds)
}

func fractionalSpeed(total Money, hs *[]models.Horse, c func(Money, Money) float64) {
	oddsPools := make([]OddsPool, len(*hs))
	totalSpeed := 0.0

	// Sum up speed.
	for _, v := range *hs {
		totalSpeed += v.Speed
	}

	previousProb := 0.0
	for i, v := range *hs {
		op := OddsPool{
			Index:     i,
			Wager:     0.0,
			Frequency: previousProb + v.Speed/totalSpeed,
		}

		previousProb = op.Frequency
		oddsPools[i] = op
	}

	rouletteDistribution(total, &oddsPools)

	for i := range *hs {
		(*hs)[i].Odds = c(oddsPools[i].Wager, total)
	}

	// fmt.Println(hs)
}

func rouletteDistribution(total Money, o *[]OddsPool) {
	sort.SliceStable(*o, func(i, j int) bool {
		return (*o)[i].Frequency < (*o)[j].Frequency
	})

	for j := 0; j < int(total); j++ {
		chance := models.R1.Float64()
		for i, v := range *o {
			if chance < v.Frequency {
				(*o)[i].Wager++
				break
			}
		}
	}

	// 	fmt.Println(o)
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

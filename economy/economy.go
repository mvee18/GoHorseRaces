package economy

import (
	"fmt"
	"horses/models"
	"sort"
)

type OddsPool struct {
	Index     int
	Wager     models.Money
	Frequency float64
}

const takePercentage = 0.15

var TotalWager = 0.0

func CalculateOdds(bet models.Money, total models.Money) float64 {
	total -= total * takePercentage
	leftOver := total - bet

	return float64(leftOver) / float64(bet)
}

func DistributeMoney(r *models.Race) {
	// Somewhere between 0-100000, but mostly around 50,000.
	TotalWager = models.R1.NormFloat64()*25000 + 50000

	fractionalSpeed(models.Money(TotalWager), &r.Horses, CalculateOdds)
}

func fractionalSpeed(total models.Money, hs *[]models.Horse, c func(models.Money, models.Money) float64) {
	oddsPools := make([]OddsPool, len(*hs))
	totalAttractiveness := 0.0

	// Sum up speed.
	for _, v := range *hs {
		totalAttractiveness += v.Attractiveness
	}

	previousProb := 0.0
	for i, v := range *hs {
		op := OddsPool{
			Index:     i,
			Wager:     0.0,
			Frequency: previousProb + v.Attractiveness/totalAttractiveness,
		}

		previousProb = op.Frequency
		oddsPools[i] = op
	}

	rouletteDistribution(total, &oddsPools)

	for i := range *hs {
		(*hs)[i].Wager = float64(oddsPools[i].Wager)
		(*hs)[i].Odds = c(oddsPools[i].Wager, total)
	}

	// fmt.Println(hs)
}

func rouletteDistribution(total models.Money, o *[]OddsPool) {
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

// < ---- User Actions and Inputs ---- >

func UpdateMoney(m *models.Money, c models.ChoiceStruct, rankings []models.Horse) {
	switch c.BetType {
	case "Win":
		winner := rankings[0]
		if c.Name == winner.Name {
			winnings := c.Bet * models.Money(winner.Odds)
			*m += winnings
			fmt.Printf("You won! The bet paid %.2f. Your total is now: %.2f\n", winnings, *m)
		} else {
			*m -= c.Bet
			fmt.Printf("You lost! Your total is now: %.2f\n", *m)
		}
	case "Place":
		winnings := calcPlaceOdds(TotalWager, rankings[0], rankings[2], c)
		*m += winnings
		if winnings > 0.0 {
			fmt.Printf("You won! The bet paid %.2f. Your total is now: %.2f\n", winnings, *m)
		} else {
			fmt.Printf("You lost! Your total is now: %.2f\n", *m)
		}
	}
}

func calcPlaceOdds(totalMoney float64, winner1, winner2 models.Horse, c models.ChoiceStruct) models.Money {
	// Total profit - amount wagered on each horse split evenly.
	houseTake := totalMoney * takePercentage
	profitPool := (totalMoney - houseTake - winner1.Wager - winner2.Wager)/2
	if profitPool < 0 {
		// Minus pool. Bettors are paid 0.05 per dollar.
		return 0.05 * c.Bet
	}
	winner1Odds := profitPool / winner1.Wager
	winner2Odds := profitPool / winner2.Wager

	// No breakage. This is a terminal game.

	if winner1.Name == c.Name {
		return models.Money(winner1Odds) * c.Bet
	} else if winner2.Name == c.Name {
		return models.Money(winner2Odds) * c.Bet
	}

	return -c.Bet
}

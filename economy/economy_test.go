package economy

import (
	"fmt"
	"horses/models"
	"math"
	"testing"
)

func TestCalculateOdds(t *testing.T) {
	t.Run("test with 900 total, 15pc take, 300 bet", func(t *testing.T) {

		got := CalculateOdds(300, 900)

		want := 1.55

		if float64(got) != want {
			fmt.Printf("wrong odds, wanted %v, got %v\n", want, got)
		}
	})
}

func TestFractionalSpeed(t *testing.T) {
	t.Run("three horse fractional speed", func(t *testing.T) {

		hs := []models.Horse{
			{
				Name:  "first",
				Speed: 1.0,
			},
			{
				Name:  "second",
				Speed: 2.0,
			},
			{
				Name:  "third",
				Speed: 5.0,
			},
		}

		fractionalSpeed(10000.0, &hs, MockCalculateOdds)

		obs := []float64{}
		for _, v := range hs {
			obs = append(obs, v.Odds)
		}

		exp := []float64{1250, 2500, 6250}

		x := ChiSquareHelper(t, exp, obs)

		if x > 5.991 {
			t.Errorf("chisq value greater than critical value of 2 df, got %v\n", x)
		}

	})
}

func ChiSquareHelper(t *testing.T, exp []float64, o []float64) float64 {
	t.Helper()

	dif := make([]float64, len(exp))

	for i, obs := range o {
		// o - e
		dif[i] = math.Pow((obs-exp[i]), 2) / exp[i]
	}

	chisq := 0.0

	for _, v := range dif {
		chisq += v
	}

	return chisq
}

// This just sets the odds to the wager so we can verify that the odds are correct with ChiSq.
func MockCalculateOdds(bet Money, total Money) float64 {
	return float64(bet)
}

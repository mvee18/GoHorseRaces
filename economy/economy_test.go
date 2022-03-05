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

		if x < 5.991 {
			t.Errorf("chisq value not significantly different from normal distribution, got chisq %v\n", x)
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
func MockCalculateOdds(bet models.Money, total models.Money) float64 {
	return float64(bet)
}

func TestUpdateMoney(t *testing.T) {
	h1 := models.Horse{
		Name: "winner",
		Odds: 2.0,
	}

	h2 := models.Horse{
		Name: "loser",
		Odds: 5.0,
	}

	t.Run("testing winner", func(t *testing.T) {
		money := models.Money(50.0)
		c := models.ChoiceStruct{
			Name:    "winner",
			Bet:     2.0,
			BetType: "Win",
		}

		rankings := []models.Horse{h1, h2}

		UpdateMoney((&money), c, rankings)

		want := models.Money(54.0)

		if money != want {
			t.Errorf("wrong money, wanted %v, got %v\n", want, money)
		}
	})

}

func TestCalcPlaceOdds(t *testing.T) {
	h1 := models.Horse{
		Name: "winner",
		Wager: 8000,
	}

	h2 := models.Horse{
		Name: "second",
		Wager: 1500,
	}

	//	h3 := models.Horse{
	//	Name: "second",
	//	Odds: 500,
	//}


	t.Run("testing with first place winner", func(t *testing.T) {
		c := models.ChoiceStruct{
			Name: "winner",
			Bet: 2.0,
			BetType: "Place",
		}

		winnings := calcPlaceOdds(20000, h1, h2, c)

		if winnings != 0.9375 {
			t.Errorf("wrong winnings, wanted 0.9375, got %v\n", winnings)
		}
	})

	t.Run("testing with second place", func(t *testing.T) {
		c := models.ChoiceStruct{
			Name: "second",
			Bet: 2.0,
			BetType: "Place",
		}

		winnings := calcPlaceOdds(20000, h1, h2, c)

		if winnings != 5 {
			t.Errorf("wrong winnings, wanted 2.5, got %v\n", winnings)
		}
	})

	t.Run("testing loser", func(t *testing.T) {
		c := models.ChoiceStruct{
			Name: "loser",
			Bet: 2.0,
			BetType: "Place",
		}

		winnings := calcPlaceOdds(20000, h1, h2, c)

		if winnings != -2.0 {
			t.Errorf("wrong loss, wanted -2, got %v\n", winnings)
		}
	})
}

func TestCalcShowOdds(t *testing.T) {
h1 := models.Horse{
		Name: "winner",
		Wager: 8000,
	}

	h2 := models.Horse{
		Name: "second",
		Wager: 1500,
	}

	h3 := models.Horse{
		Name: "third",
		Wager: 500,
	}

	hs := []models.Horse{h1, h2, h3}
	t.Run("testing third place winner", func(t *testing.T) {
	c := models.ChoiceStruct{
			Name: "third",
			Bet: 2.0,
			BetType: "Show",
		}

		winnings := fmt.Sprintf("%.2f", calcShowOdds(15000, hs, c))
		want := "3.67"
		if winnings != want {
			t.Errorf("wrong winnings, wanted %v, got %v\n", want, winnings)
		}

	})

	t.Run("testing loser", func(t *testing.T) {
		c := models.ChoiceStruct{
				Name: "loser",
				Bet: 2.0,
				BetType: "Show",
			}

			winnings := fmt.Sprintf("%.2f", calcShowOdds(15000, hs, c))
			want := "-2.00"
			if winnings != want {
				t.Errorf("wrong winnings, wanted %v, got %v\n", want, winnings)
			}

	})
}

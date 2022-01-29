package economy

import (
	"fmt"
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

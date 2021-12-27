package main

import (
	"fmt"
	"testing"
)

func TestGenerateRace(t *testing.T) {
	t.Run("testing race generation", func(t *testing.T) {
		r := GenerateRace(2, "fast", 0.75)

		if len(r.Horses) != 2 || r.Weather != "fast" || r.Length != 0.75 {
			t.Errorf("wrong parameters in horse race generation")
		}

		fmt.Printf("%#v\n", r.Horses)
	})

	t.Run("Test normalize speed", func(t *testing.T) {
		h1 := Horse{
			Name:  "tester1",
			Speed: 8.123,
		}

		h2 := Horse{
			Name:  "tester2",
			Speed: 123.67,
		}

		horses := []Horse{h1, h2}

		normalizeSpeed(horses)

		h1want := h1.Speed / h1.Speed
		h2want := h2.Speed / h1.Speed

		if horses[0].Speed != h1want {
			t.Errorf("wrong h1 speed, wanted %v got %v\n", h1want, horses[0].Speed)
		}

		if horses[1].Speed != h2want {
			t.Errorf("wrong h2 speed, wanted %v got %v\n", h2want, horses[1].Speed)
		}
	})

}

func TestShowRace(t *testing.T) {
	t.Run("testing race", func(t *testing.T) {
		ShowRace()
	})
}

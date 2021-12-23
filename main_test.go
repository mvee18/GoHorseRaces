package main

import (
	"fmt"
	"testing"
)

func TestGenerateHorse(t *testing.T) {
	t.Run("testing horse generation", func(t *testing.T) {
		h := GenerateHorse()

		fmt.Printf("%#v\n", h)

		if h.Name == "" {
			t.Errorf("wanted non nil horse name\n")
		}
	})
}

func TestNormalDistribution(t *testing.T) {
	t.Run("normal value?", func(t *testing.T) {
		normalDistribution()
	})
}

package main

import (
	"testing"
)

func TestNormalDistribution(t *testing.T) {
	t.Run("normal value?", func(t *testing.T) {
		normalDistribution()
	})
}

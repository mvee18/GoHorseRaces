package models

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
func TestGenerateRecord(t *testing.T) {
	t.Run("testing record generation", func(t *testing.T) {
		h := Horse{
			Name:   "test",
			Speed:  10,
			Odds:   0,
			Record: Record{},
			Jockey: *GenerateJockey(),
		}

		h.generateRecord()

		fmt.Printf("the record is %v\n", h.Record)
	})
}

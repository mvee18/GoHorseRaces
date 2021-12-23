package main

import "sort"

type Race struct {
	Horses  []Horse
	Length  float64
	Weather string
}

func GenerateRace(n int, w string, l float64) Race {
	h := []Horse{}

	for i := 0; i < n; i++ {
		h = append(h, GenerateHorse())
	}

	r := Race{
		Horses:  h,
		Weather: w,
		Length:  l,
	}

	return r
}

func normalizeSpeed(h []Horse) {
	// We need to sort the horses by speed to find the minimum value.
	sort.SliceStable(h, func(i, j int) bool {
		return h[i].Speed < h[j].Speed
	})

	minimumSpeed := h[0].Speed

	for i := 0; i < len(h); i++ {
		h[i].Speed = h[i].Speed / minimumSpeed
	}
}

package main

import (
	"fmt"
	"runtime"
	"sort"
	"sync"
	"time"

	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
)

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

	sort.SliceStable(h, func(i, j int) bool {
		return h[i].Speed < h[j].Speed
	})

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

func makeRaceBars(hs []Horse) {
	// start the progress bars in go routines
	var wg sync.WaitGroup

	p := mpb.New(mpb.WithWaitGroup(&wg))
	total, numBars := 100, len(hs)
	// start := time.Now()

	wg.Add(numBars)

	for i := 0; i < numBars; i++ {
		name := fmt.Sprintf("Horse %d:", i)
		bar := p.AddBar(int64(total),
			mpb.PrependDecorators(
				// simple name decorator
				decor.Name(name),
				// decor.DSyncWidth bit enables column width synchronization
				decor.Percentage(decor.WCSyncSpace),
			),
			mpb.AppendDecorators(
				// replace ETA decorator with "done" message, OnComplete event
				decor.OnComplete(
					// ETA decorator with ewma age of 60
					decor.EwmaETA(decor.ET_STYLE_GO, 60, decor.WCSyncWidth), "done",
				),
			),
		)

		go addHorseBar(i, bar, total, hs[i], &wg)
	}

	/*
		for i, h := range hs {
			wg.Add(1)
			go horseRacer(i, h, &wg)
		}
	*/

	wg.Wait()
	time.Sleep(time.Millisecond * 100)
}

func addHorseBar(index int, bar *mpb.Bar, barTot int, h Horse, wg *sync.WaitGroup) {
	defer wg.Done()

	i := 0
	for i <= barTot {
		if determineIfProceed(&h) {
			bar.Increment()
			time.Sleep(time.Millisecond * 1)
			i++
		} else {
			time.Sleep(time.Millisecond * 1)
		}
	}

	// fmt.Printf("Horse %d, %s, has finished with speed %.2f!\n", index, h.Name, h.Speed)

}

func horseRacer(index int, h Horse, wg *sync.WaitGroup) {
	defer wg.Done()

	count := 0
	time.Sleep(time.Millisecond * 100)
	for count < 10 {
		if determineIfProceed(&h) {
			count += 1
		} else {
			continue
		}
	}

	fmt.Printf("Horse %d, %s, has finished with speed %.2f!\n", index, h.Name, h.Speed)
}

func determineIfProceed(h *Horse) bool {
	// The maximum value of the horse speed is 20, since the maximum horse
	// value is 10, and the maximum horse jocket benefit is 100%.
	roll := 0.0 + r1.Float64()*(20.0-0.0)

	if roll < h.Speed {
		return true
	}

	return false
}

func ShowRace() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	r := GenerateRace(6, "fast", 0.75)

	makeRaceBars(r.Horses)
}

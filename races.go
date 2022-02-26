package main

import (
	"fmt"
	"horses/economy"
	"horses/models"
	"horses/ui"
	"math/rand"
	"os"
	"runtime"
	"sort"
	"sync"
	"text/tabwriter"
	"time"

	"github.com/cheynewallace/tabby"
	"github.com/kyokomi/emoji/v2"
	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
)

func GenerateRace(n int, w string, l float64) models.Race {
	h := []models.Horse{}

	for i := 0; i < n; i++ {
		h = append(h, models.GenerateHorse())
	}

	r := models.Race{
		Horses:  h,
		Weather: w,
		Length:  l,
	}

	sort.SliceStable(h, func(i, j int) bool {
		return h[i].Speed < h[j].Speed
	})

	economy.DistributeMoney(&r)

	return r
}

func makeRaceBars(hs []models.Horse) []models.Horse {
	// start the progress bars in go routines
	results := make(chan int, len(hs))
	var wg sync.WaitGroup

	p := mpb.New(mpb.WithWaitGroup(&wg))
	total, numBars := 100, len(hs)
	// start := time.Now()

	// Without the shuffling, the order of the horses is always least to greatest.
	Shuffle(hs)

	wg.Add(numBars)
	emojiLine := emoji.Sprint("[=:horse:-]<+")

	for i := 0; i < numBars; i++ {
		name := fmt.Sprintf("%s", hs[i].Name)
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
			mpb.BarStyle(emojiLine),
			//mpb.BarStyle("[=H-]<+"),
		)


		go addHorseBar(i, bar, total, hs[i], &wg, results)
	}

	/*
		for i, h := range hs {
			wg.Add(1)
			go horseRacer(i, h, &wg)
		}
	*/

	p.Wait()

	time.Sleep(time.Millisecond * 200)

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	rankings := printResults(w, hs, results)
	w.Flush()

	return rankings
}

func printResults(w *tabwriter.Writer, hs []models.Horse, results chan int) []models.Horse {
	t := tabby.NewCustom(w)

	fmt.Printf("The standings are as follows:\n")
	t.AddHeader("PLACEMENT", "NAME", "ODDS")

	rankings := make([]models.Horse, len(hs))

	for i := 0; i < len(hs); i++ {
		select {
		case placement := <-results:
			h := hs[placement]
			rankings[i] = hs[i]
			t.AddLine(i, h.Name, fmt.Sprintf("%.2f", h.Odds))
		}
	}

	t.Print()

	return rankings
}

func Shuffle(vals []models.Horse) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(vals) > 0 {
		n := len(vals)
		randIndex := r.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
		vals = vals[:n-1]
	}
}

func addHorseBar(index int, bar *mpb.Bar, barTot int, h models.Horse, wg *sync.WaitGroup, out chan<- int) {
	defer wg.Done()

	i := 0
	for i <= barTot {
		if determineIfProceed(&h) {
			bar.Increment()
			time.Sleep(time.Millisecond * 3)
			i++
		} else {
			time.Sleep(time.Millisecond * 3)
		}
	}

	// Write the
	out <- index

	// fmt.Printf("Horse %d, %s, has finished with speed %.2f!\n", index, h.Name, h.Speed)

}

func horseRacer(index int, h models.Horse, wg *sync.WaitGroup) {
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

func determineIfProceed(h *models.Horse) bool {
	// The maximum value of the horse speed is 20, since the maximum horse
	// value is 10, and the maximum horse jocket benefit is 100%.
	roll := 0.0 + r1.Float64()*(20.0-0.0)

	if roll < h.Speed {
		return true
	}

	return false
}

func ShowRace(m *models.Money) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	r := GenerateRace(6, "fast", 0.75)

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	printRace(w, r.Horses)

	choiceStruct, err := ui.ShowList(r.Horses)
	if err != nil {
		panic(err)
	}

	rankings := makeRaceBars(r.Horses)
	winner := rankings[0]

	fmt.Printf("The winner is %s!\n", rankings[0].Name)

	if choiceStruct.Name == winner.Name {
		economy.UpdateMoney(m, choiceStruct.BetType, choiceStruct.Name, rankings)
	} else {
		*m = *m - choiceStruct.Bet
	}

	fmt.Printf("Your new money is %.2f\n", *m)
}



func printRace(w *tabwriter.Writer, hs []models.Horse) {
	t := tabby.NewCustom(w)

	sort.SliceStable(hs, func(i, j int) bool {
		return hs[i].Odds < hs[j].Odds
	})

	fmt.Printf("The standings are as follows:\n")
	t.AddHeader("NAME", "ODDS", "WINS", "PLACES", "SHOWS")

	for i := 0; i < len(hs); i++ {
		h := hs[i]
		t.AddLine(h.Name, fmt.Sprintf("%.2f", h.Odds), h.Record.Wins, h.Record.Places, h.Record.Shows)
	}

	t.Print()

}

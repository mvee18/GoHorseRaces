package main

import (
	"fmt"
	"horses/economy"
	"math/rand"
	"time"
)

const LineBreak = "\n"

var r1 *rand.Rand

func init() {
	r1 = rand.New(rand.NewSource(time.Now().UnixNano() + 2561))
}

func normalDistribution() {
	n := r1.NormFloat64()
	fmt.Println(n)
}

func main() {
	/*
		fmt.Println("What would you like to bet on?")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("There was an error reading your input. Please try again.")
		}

		input = strings.TrimSuffix(input, LineBreak)

		fmt.Println(input)
	*/

	/*
		uiprogress.Start()            // start rendering
		bar := uiprogress.AddBar(100) // Add a new bar

		// optionally, append and prepend completion and elapsed time
		bar.AppendCompleted()
		bar.PrependElapsed()

		for bar.Incr() {
			time.Sleep(time.Millisecond * 20)
		}
	*/
	m := economy.Money(200.0)

	for m > 0.0 {
		ShowRace(&m)
	}
}

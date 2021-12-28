package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/goombaio/namegenerator"
)

var NameGenerator namegenerator.Generator
var r1 *rand.Rand

func init() {
	seed := time.Now().UTC().UnixNano()
	NameGenerator = namegenerator.NewNameGenerator(seed)

	r1 = rand.New(rand.NewSource(time.Now().UnixNano()))
}

const LineBreak = "\n"

type Horse struct {
	Name  string
	Speed float64
	Odds  float64
	Jockey
	Record
}

func GenerateHorse() Horse {
	h := Horse{}
	h.Name = generateName()

	h.Jockey = *GenerateJockey()

	h.Speed = generateSpeed() * h.Jockey.Benefit

	return h
}

// Probably could improve the name generator to use "real" names for the jockeys.
func generateName() string {
	return NameGenerator.Generate()
}

func generateSpeed() float64 {
	return 0.0 + r1.Float64()*(10.0-0.0)
}

type Record struct {
	Wins    int
	Shows   int
	Places  int
	Fourths int
}

type Jockey struct {
	Name    string
	Benefit float64
	Record
}

func GenerateJockey() *Jockey {
	j := Jockey{}

	j.Name = generateName()

	/*
		Let's just say the jockey confers a percentage based bonus to the horse.
		This would be between 0 and 1 but we will multiply so add one to be a
		multiplicative bonus.
	*/
	j.Benefit = r1.Float64() + 1

	return &j
}

func normalDistribution() {
	n := r1.NormFloat64()
	fmt.Println(n)
}

// func main() {
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

//	ShowRace()
// }

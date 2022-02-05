package models

import (
	"math/rand"
	"time"

	"github.com/goombaio/namegenerator"
)

var NameGenerator namegenerator.Generator
var R1 *rand.Rand

const MaxSpeed = 10.0
const MinSpeed = 1.0
const SpeedFraction = 0.40

func init() {
	seed := time.Now().UTC().UnixNano()
	NameGenerator = namegenerator.NewNameGenerator(seed)

	R1 = rand.New(rand.NewSource(time.Now().UnixNano()))
}

type Horse struct {
	Name           string
	Attractiveness float64
	Speed          float64
	Odds           float64
	Jockey
	Record
}

func (h *Horse) generateRecord() {
	// A random number of a random number within 10 for the total of races
	// the horse has participated in.
	total := R1.Intn(10) + 5

	wins := R1.Intn(total)

	// fmt.Printf("tot - win: %d\n", total-wins)
	place := R1.Intn(total - wins)

	shows := R1.Intn(total - wins - place)

	h.Record = Record{
		Wins:   wins,
		Places: place,
		Shows:  shows,
		Total:  total,
	}
}

type Record struct {
	Wins   int
	Shows  int
	Places int
	Total  int
}

type Jockey struct {
	Name    string
	Benefit float64
	Record
}

type Race struct {
	Horses  []Horse
	Length  float64
	Weather string
}

func GenerateJockey() *Jockey {
	j := Jockey{}

	j.Name = generateName()

	/*
		Let's just say the jockey confers a percentage based bonus to the horse.
		This would be between 0 and 1 but we will multiply so add one to be a
		multiplicative bonus.
	*/
	j.Benefit = R1.Float64() + 1

	return &j
}

func GenerateHorse() Horse {
	h := Horse{
		Name:           generateName(),
		Attractiveness: 0.0,
		Speed:          0.0,
		Odds:           0.0,
		Jockey:         *GenerateJockey(),
	}

	h.Speed = generateSpeed() * h.Jockey.Benefit
	h.calculateAttractiveness()
	h.generateRecord()

	return h
}

func (h *Horse) calculateAttractiveness() {
	random := 0.0 + R1.Float64()*(11.0-0.0)
	h.Attractiveness = h.Speed*0.50 + random
}

// Probably could improve the name generator to use "real" names for the jockeys.
func generateName() string {
	return NameGenerator.Generate()
}

func generateSpeed() float64 {
	return MinSpeed + R1.Float64()*(MaxSpeed-MinSpeed)
}

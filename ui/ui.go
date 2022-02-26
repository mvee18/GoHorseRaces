package ui

import (
	"fmt"
	"os"
	"strconv"

	"github.com/erikgeiser/promptkit/selection"
	"github.com/erikgeiser/promptkit/textinput"
	"horses/models"
)

var BetList = []string{"Win", "Place", "Show"}

func ShowList(hs []models.Horse) (models.ChoiceStruct, error) {
	choices := getHorseNames(hs)
	c := horseList(choices)

	verified := false
	for !verified {
		verified = verifyChoice(c)
		if verified {
			break
		} else {
			c = horseList(choices)
		}
	}

	btype := betType()

	bet, err := getBet(c)
	if err != nil {
		return models.ChoiceStruct{}, err
	}

	return models.ChoiceStruct{Name: c.String, Bet: bet, BetType: btype}, nil
}

func getHorseNames(h []models.Horse) []string {
	names := make([]string, len(h))

	for i, v := range h {
		names[i] = v.Name
	}

	return names
}

func horseList(c []string) *selection.Choice {
	sp := selection.New("Which horse do you bet on?",
		selection.Choices(c))
	sp.PageSize = len(c)

	choice, err := sp.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}

	return choice
}

func verifyChoice(s *selection.Choice) bool {
	sp := selection.New(fmt.Sprintf("You have chosen, %s, is that correct?", s.String),
		selection.Choices([]string{"Yes", "No"}))
	sp.PageSize = 2

	choice, err := sp.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}

	if choice.String == "Yes" {
		return true
	} else {
		return false
	}
}

func betType() string {
	sp := selection.New("Which type of bet do you wish to place?",
		selection.Choices(BetList))
	sp.PageSize = len(BetList)

	choice, err := sp.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}

	return choice.String

}

func getBet(choice *selection.Choice) (models.Money, error) {
	input := textinput.New(fmt.Sprintf("How much do you want to bet on %s?", choice.String))
	input.Placeholder = "You cannot bet $0."

	resp, err := input.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}

	// do something with the result
	bet, err := strconv.ParseFloat(resp, 64)
	if err != nil {
		return 0.0, err
	}

	return models.Money(bet), nil
}

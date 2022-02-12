package ui

import (
	"fmt"
	"horses/economy"
	"os"
	"strconv"

	"github.com/erikgeiser/promptkit/selection"
	"github.com/erikgeiser/promptkit/textinput"
)

type ChoiceStruct struct {
	Name string
	Bet  economy.Money
}

func ShowList(choices []string) (ChoiceStruct, error) {
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

	bet, err := getBet(c)
	if err != nil {
		return ChoiceStruct{}, err
	}

	return ChoiceStruct{Name: c.String, Bet: bet}, nil
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

func getBet(choice *selection.Choice) (economy.Money, error) {
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

	return economy.Money(bet), nil
}

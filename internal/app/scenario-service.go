package app

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/charmbracelet/huh"
)

// Creates the scenario background
func InitStory(state *GameState) *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("What's your name?").
				Value(&state.Name),
			huh.NewInput().
				Title("What's your backstory?").
				Value(&state.Description),
			huh.NewSelect[int]().
				Title("What's your age?").
				Options(
					huh.NewOption("Young", 15),
					huh.NewOption("Middle-aged", 25),
					huh.NewOption("Elderly", 60),
				).
				Value(&state.Age),
		),
	)
}

func RunChoiceBuilder(state *GameState) *huh.Form {
	// Fixed: Use a constant value or meaningful options
	burgerOptions := []huh.Option[int]{
		huh.NewOption("Healthy Burger (5 HP)", 5),
		huh.NewOption("Cheeseburger (3 HP)", 3),
		huh.NewOption("Mystery Burger", rand.Intn(10)+1),
	}

	return huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().
				Title("Choose your burger").
				Options(burgerOptions...).
				Value(&state.Health),
		),
	)
}

func RunChoiceBuilderN(state *GameState, nChoices int) (*huh.Form, error) {
	if nChoices <= 0 {
		return nil, errors.New("There cannot be less than 1 choice")
	}

	randomOptions := make([]huh.Option[string], nChoices)
	for i := range nChoices {
		scenarioToReturn := CreateRandomScenarioChoice()
		option := huh.NewOption(scenarioToReturn, scenarioToReturn)
		randomOptions[i] = option
	}

	return huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose your next move").
				Options(randomOptions...).
				Value(&state.NextSteps),
		),
	), nil
}

func CreateRandomScenarioChoice() string {
	timeNow := time.Now()

	// Generate a random delay between 1-5 seconds
	delay := time.Second * time.Duration(1+rand.Intn(4)) // 1-5 seconds
	time.Sleep(delay)

	return fmt.Sprintf("Random scenario %v", time.Since(timeNow))
}

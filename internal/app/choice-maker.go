package app

import (
	"math/rand"

	"github.com/charmbracelet/huh"
)

func RunChoiceBuilder(state *GameState) *huh.Form {
	// Fixed: Use a constant value or meaningful options
	burgerOptions := []huh.Option[int]{
		huh.NewOption("Healthy Burger (5 HP)", 5),
		huh.NewOption("Cheeseburger (3 HP)", 3),
		huh.NewOption("Mystery Burger", rand.Intn(10)+1),
	}

	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("What's your name?").
				Value(&state.Name),
		),
		huh.NewGroup(
			huh.NewSelect[int]().
				Title("Choose your burger").
				Options(burgerOptions...).
				Value(&state.Health),
		),
	)
}

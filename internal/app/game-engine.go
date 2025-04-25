package app

import (
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
)

type GameEngine struct {}

type GameState struct {
	Name   string
	Health int
}

func (g *GameEngine) Run() {
	state := &GameState{}
	var form *huh.Form
	for {
		form = RunChoiceBuilder(state)
		err := form.Run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("\n--- Game State ---")
		fmt.Printf("Name: %s\nHealth: %d\n", state.Name, state.Health)
		fmt.Printf("Full state: %+v\n", *state) // %+v shows field names
	}
}

package app

import "fmt"

type GameState struct {
	Name        string
	Description string
	Age         int
	Health      int
	AttackDmg   int
	Defence     int
	NextSteps   string
}

func (state *GameState) AnnounceGameState() {
	fmt.Println("\n--- Game State ---")
	fmt.Printf("Name: %s\nHealth: %d\n", state.Name, state.Health)
	fmt.Printf("Full state: %+v\n", *state) // %+v shows field names
}

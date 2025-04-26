package app

import "fmt"

type GameState struct {
	PlayerDetails
	EventState
}

type EventState struct {
	// The Player's Previous Steps
	AllSteps []string
	// The Narration's Question
	Narration string
	// The Player's Next Steps
	NextSteps string
}

type PlayerDetails struct {
	Name        string
	Description string
	Age         int
	Health      int
	AttackDmg   int
	Defence     int
}

func (state *GameState) AnnounceGameState() {
	fmt.Println("\n--- Game State ---")
	fmt.Printf("Name: %s\nHealth: %d\n", state.Name, state.Health)
	fmt.Printf("Full state: %+v\n", *state) // %+v shows field names
}

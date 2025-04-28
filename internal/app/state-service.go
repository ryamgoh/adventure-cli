package app

import (
	"fmt"
	"math/rand"
)

type GameState struct {
	PlayerDetails
	EventState
}

type Event struct {
	Role        Role
	Description string
}

// Event state: history and current events
type EventState struct {
	EventHistory []Event // All past events, User and Narrator
	Narration    Event   // Current narration, must be Narrator event
	NextSteps    Event   // Next user step, must be User event
}

// Add event to history log
func (es *EventState) AddEventToHistory(e Event) {
	es.EventHistory = append(es.EventHistory, e)
}

func (es *EventState) AddEventsToHistory() {
	es.AddEventToHistory(es.Narration)
	es.AddEventToHistory(es.NextSteps)
}

type Role int

const (
	User Role = iota + 1 // EnumIndex = 1
	Narrator
)

func (r Role) String() string {
	return [...]string{"User", "Narrator"}[r-1]
}

func (r Role) EnumIndex() int {
	return int(r)
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

func (state *GameState) CreateSessionOnUser() {
	state.NextSteps.Description = fmt.Sprintf(
		`
		The User Details
		Name: %v
		Age: %v
		Description: %v
		Health: %v
		Attack: %v
		Defence: %v
		`,
		state.Name,
		state.Age,
		state.Description,
		state.Health,
		state.AttackDmg,
		state.Defence,
	)
}

func (state *GameState) GetNextNarration() {
	state.Narration.Description = fmt.Sprintf(
		"This narration description is entirely arbitrary: Narration %v", rand.Intn(10),
	)
}

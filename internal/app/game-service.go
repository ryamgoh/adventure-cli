package app

import (
	"log"

	"github.com/charmbracelet/huh"
)

type GameEngine struct {
	State *GameState
}

func (g *GameEngine) initGameEngine() {
	g.State = &GameState{}
	g.State.EventHistory = make([]Event, 0, 100)
	g.State.NextSteps = Event{
		Role:        User,
		Description: "",
	}
	g.State.Narration = Event{
		Role:        Narrator,
		Description: "",
	}
}

func (g *GameEngine) Run() {
	g.initGameEngine()
	var state = g.State
	var error error
	var scenario *huh.Form
	runScenario(InitStory(state))
	state.CreateSessionOnUser()
	state.AddEventToHistory(state.NextSteps)
	for {
		scenario, error = RunChoiceBuilderN(state, 4)
		if error != nil {
			break
		}
		runScenario(scenario)
		state.AddAllEvents()
		state.AnnounceGameState()
	}
	log.Fatal("Broken out of loop")
}

func runScenario(scenario *huh.Form) {
	err := scenario.Run()
	if err != nil {
		log.Fatal(err)
	}
}

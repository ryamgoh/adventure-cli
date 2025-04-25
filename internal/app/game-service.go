package app

import (
	"log"

	"github.com/charmbracelet/huh"
)

type GameEngine struct{
	State *GameState
}

func (g *GameEngine) initGameEngine() {
	g.State = &GameState{}
}

func (g *GameEngine) Run() {
	g.initGameEngine()
	var state = g.State
	var error error
	var scenario *huh.Form
	scenario = InitStory(state)
	runScenario(scenario)
	for {
		// scenario = RunChoiceBuilder(state)
		scenario, error = RunChoiceBuilderN(state, 4)
		if (error != nil) {
			break
		}
		runScenario(scenario)
		state.AnnounceGameState()
	}
}

func runScenario(scenario *huh.Form) {
	err := scenario.Run()
	if err != nil {
		log.Fatal(err)
	}
}

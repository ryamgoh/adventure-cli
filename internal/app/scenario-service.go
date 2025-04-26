package app

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
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
		huh.NewGroup(
			huh.NewSelect[int]().
				Title("Whats your attack damage?").
				Options(
					huh.NewOption("Weak", 20 + rand.Intn(8)),
					huh.NewOption("Normal", 30 + rand.Intn(5)),
					huh.NewOption("Strong", 45 + rand.Intn(8)),
				).
				Value(&state.AttackDmg),
			huh.NewSelect[int]().
				Title("Whats your defence?").
				Options(
					huh.NewOption("Weak", 20 + rand.Intn(8)),
					huh.NewOption("Normal", 30 + rand.Intn(5)),
					huh.NewOption("Strong", 45 + rand.Intn(8)),
				).
				Value(&state.Defence),
			huh.NewSelect[int]().
				Title("Whats your health?").
				Options(
					huh.NewOption("Weak", 20 + rand.Intn(8)),
					huh.NewOption("Normal", 30 + rand.Intn(5)),
					huh.NewOption("Strong", 45 + rand.Intn(8)),
				).
				Value(&state.Health),
		),
	)
}

func RunChoiceBuilderN(state *GameState, nChoices int) (*huh.Form, error) {
	if nChoices <= 0 {
		return nil, errors.New("there cannot be less than 1 choice")
	}

	var wg sync.WaitGroup
	var mu sync.RWMutex
	randomOptions := make([]huh.Option[string], 0, nChoices)

	// Launch goroutines to generate scenarios concurrently
	for range nChoices {
		wg.Add(1)
		go func() {
			defer wg.Done()
			scenario := CreateRandomScenarioChoice()
			mu.Lock()
			randomOptions = append(randomOptions, huh.NewOption(scenario, scenario))
			mu.Unlock()
		}()
	}

	// Spinner is only for waiting, so just wrap wg.Wait()
	err := spinner.New().
		Title("Generating new choices...").
		Action(func() {
			wg.Wait() // spinner displays while we're waiting
		}).
		Run()
	if err != nil {
		return nil, err
	}

	return huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title(state.Narration.Description).
				Options(randomOptions...).
				Value(&state.NextSteps.Description),
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

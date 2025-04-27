package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms/openai"
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
					huh.NewOption("Weak", 20+rand.Intn(8)),
					huh.NewOption("Normal", 30+rand.Intn(5)),
					huh.NewOption("Strong", 45+rand.Intn(8)),
				).
				Value(&state.AttackDmg),
			huh.NewSelect[int]().
				Title("Whats your defence?").
				Options(
					huh.NewOption("Weak", 20+rand.Intn(8)),
					huh.NewOption("Normal", 30+rand.Intn(5)),
					huh.NewOption("Strong", 45+rand.Intn(8)),
				).
				Value(&state.Defence),
			huh.NewSelect[int]().
				Title("Whats your health?").
				Options(
					huh.NewOption("Weak", 20+rand.Intn(8)),
					huh.NewOption("Normal", 30+rand.Intn(5)),
					huh.NewOption("Strong", 45+rand.Intn(8)),
				).
				Value(&state.Health),
		),
	)
}

// LLMGameResponse holds the expected response format from the LLM
type LLMGameResponse struct {
	Narration string   `json:"narration"`
	Options   []string `json:"options"`
}

// CallOpenAILLM gets narration and 4 options from OpenAI given history, narration, and user option
func CallOpenAILLM(
	ctx context.Context,
	llm *openai.LLM,
	eventHistory []Event,
	lastUserOption string,
	currentNarration string,
) (
	narration string, 
	options []string, 
	err error) {

	var historySb strings.Builder
	for _, evt := range eventHistory {
		historySb.WriteString(fmt.Sprintf("%s: %s\n", evt.Role, evt.Description))
	}

	systemPrompt := `You are a creative interactive story engine. 
		Given the story history, the last narration, and the last user action, return ONLY a valid JSON object of this form:
			{
				"narration": "...",
				"options": [
					"...",
					"...",
					"...",
					"..."
				]
			}
		`

	fullPrompt := fmt.Sprintf(`%s
		STORY HISTORY:
		%s

		LAST NARRATION: %s

		LAST USER ACTION: %s

		Strictly reply with one valid JSON object using the above schema, no extra commentary.`,
		systemPrompt, historySb.String(), currentNarration, lastUserOption)

	log.Println("Full Prompt", fullPrompt)

	resp, err := llm.Call(ctx, fullPrompt)
	if err != nil {
		return "", nil, fmt.Errorf("LLM call failed: %w", err)
	}

	// Sometimes the model returns preamble, so pick the first JSON found
	jsonStart := strings.Index(resp, "{")
	if jsonStart < 0 {
		return "", nil, fmt.Errorf("no JSON found in LLM response: %q", resp)
	}
	trimmed := resp[jsonStart:]

	log.Println("Trimmed", trimmed)

	var parsed LLMGameResponse
	if err := json.Unmarshal([]byte(trimmed), &parsed); err != nil {
		return "", nil, fmt.Errorf("invalid LLM JSON: %w\nRaw response: %s", err, resp)
	}
	if len(parsed.Options) != 4 {
		return "", nil, fmt.Errorf("expected 4 options, got %d, response: %s", len(parsed.Options), resp)
	}
	return parsed.Narration, parsed.Options, nil
}

// Usage in your choice builder:
func RunChoiceBuilderN(state *GameState, nChoices int) (*huh.Form, error) {
	if nChoices != 4 {
		return nil, fmt.Errorf("OpenAI LLM expects exactly 4 options for this prompt style")
	}

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Could not load .env file - using system environment variables")
	}

	// Get API key from environment
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY environment variable not set")
	}

	ctx := context.Background()
	// Initialize OpenAI client
	openaiLLM, err := openai.New(openai.WithToken(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize OpenAI client: %w", err)
	}

	narration, options, err := CallOpenAILLM(
		ctx, openaiLLM,
		state.EventHistory,
		state.NextSteps.Description,
		state.Narration.Description,
	)
	if err != nil {
		return nil, err
	}
	log.Println("New Narration", narration)
	log.Println("New Options", options)


	state.Narration.Description = narration
	randomOptions := make([]huh.Option[string], 0, 4)
	for _, opt := range options {
		randomOptions = append(randomOptions, huh.NewOption(opt, opt))
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


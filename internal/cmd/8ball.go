package cmd

import (
	"github.com/akl-infra/bot/internal/cmd/cmd_core"
	"golang.org/x/exp/rand"
	"time"
)

type EightBall struct{ cmd_core.DefaultCommand }

func (cmd EightBall) Exec(_ string, _ uint64) (string, error) {
	rand.Seed(uint64(time.Now().UnixNano()))
	responses := []string{
		"Yes", "Count on it",
		"No doubt",
		"Absolutely", "Very likely",
		"Maybe", "Perhaps",
		"No", "No chance", "Unlikely",
		"Doubtful", "Probably not",
	}
	return responses[rand.Intn(len(responses))], nil
}

func (cmd EightBall) Use() string {
	return "!akl 8ball"
}

func (cmd EightBall) Desc() string {
	return "8ball"
}

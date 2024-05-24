package cmd

import (
	"github.com/akl-infra/bot/internal/cmd/cmd_core"
)

type Help struct{ cmd_core.DefaultCommand }

func (cmd Help) Exec(_ string, _ uint64) (string, error) {
	return "AKL AKL AKL!", nil
}

func (cmd Help) Use() string {
	return "!akl help [command]"
}

func (cmd Help) Desc() string {
	return "help"
}

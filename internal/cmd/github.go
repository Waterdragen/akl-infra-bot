package cmd

import "github.com/akl-infra/bot/internal/cmd/cmd_core"

type Github struct{ cmd_core.DefaultCommand }

func (cmd Github) Exec(_ string, _ uint64) (string, error) {
	return "<https://github.com/akl-infra>", nil
}

func (cmd Github) Use() string {
	return "!akl github"
}

func (cmd Github) Desc() string {
	return "github"
}

package handler

import (
	"errors"
	"github.com/akl-infra/bot/internal/cmd"
	"github.com/akl-infra/bot/internal/cmd/cmd_core"
)

type CommandFunc func(string, []string) (string, error)

var Commands = map[string]cmd_core.Commandable{
	"8ball":  cmd.EightBall{},
	"gh":     cmd.Github{},
	"github": cmd.Github{},
	"help":   cmd.Help{},
	"ping":   cmd.Ping{},
	"list":   cmd.List{},
	"view":   cmd.View{},
}

func Dispatch(command string, arg string, uuid uint64) (string, error) {
	if commandObject, found := Commands[command]; !found {
		return "Command not found", errors.New("Command not found")
	} else {
		return commandObject.Exec(arg, uuid)
	}
}

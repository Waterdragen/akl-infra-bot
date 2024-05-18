package handler

import (
	"errors"

	"github.com/akl-infra/bot/internal/cmd"
)

type CommandFunc func(string, []string) (string, error)

var Commands = map[string]CommandFunc{
	"ping": cmd.Ping,
	"help": cmd.Help,
	"list": cmd.List,
	"view": cmd.View,
}

func Dispatch(command string, args []string) (string, error) {
	if f, ok := Commands[command]; !ok {
		return "Command not found", errors.New("Command not found")
	} else {
		return f(command, args)
	}
}

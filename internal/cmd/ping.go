package cmd

import (
	"fmt"
	"github.com/akl-infra/bot/internal/cmd/cmd_core"
	"github.com/akl-infra/bot/internal/util/parser"
	"strings"
)

type Ping struct{ cmd_core.DefaultCommand }

func (cmd Ping) Exec(arg string, _ uint64) (string, error) {
	args := parser.ParseQuotedArgs(arg)
	return fmt.Sprintf("ping: [%s]", strings.Join(args, ",")), nil
}

func (cmd Ping) Use() string {
	return "!akl ping"
}

func (cmd Ping) Desc() string {
	return "ping"
}

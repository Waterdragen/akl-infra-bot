package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/akl-infra/bot/internal/cmd/cmd_core"
	"github.com/akl-infra/bot/internal/util/parser"
	"io"
	"strings"

	"github.com/akl-infra/bot/internal/api"
	"github.com/akl-infra/slf/v2"
)

type View struct{ cmd_core.DefaultCommand }

func (cmd View) Exec(arg string, _ uint64) (string, error) {
	args := parser.ParseArgs(arg)
	if len(args) == 0 {
		return "No layout specified", nil
	}

	layoutName := fmt.Sprintf("layout/%s", args[0])

	res, err := api.Get(layoutName)
	if err != nil {
		return "Error contacting API", err
	}
	defer res.Body.Close()

	var body []byte
	var layout slf.Layout

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return "Error reading response from API", err
	}
	if err := json.Unmarshal(body, &layout); err != nil {
		return "Error parsing response from API", err
	}

	var keys [][]string
	var sb strings.Builder

	fmt.Fprintf(&sb, "%s (%s)\n", layout.Name, layout.Author)

	for _, key := range layout.Keys {
		if len(keys) <= int(key.Row) {
			keys = append(keys, []string{})
		}
		keys[key.Row] = append(keys[key.Row], key.Char)
	}
	for _, row := range keys {
		fmt.Fprintf(&sb, "%s\n", strings.Join(row, " "))
	}

	return sb.String(), nil
}

func (cmd View) Use() string {
	return "!akl view"
}

func (cmd View) Desc() string {
	return "view layouts"
}

package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/akl-infra/bot/internal/cmd/cmd_core"
	"io"
	"strings"

	"github.com/akl-infra/bot/internal/api"
)

type List struct{ cmd_core.DefaultCommand }

func (cmd List) Exec(_ string, _ uint64) (string, error) {
	res, err := api.Get("layouts")
	if err != nil {
		return "Error contacting API", err
	}
	defer res.Body.Close()
	var body []byte
	var layouts []string
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return "Error reading response from API", err
	}
	if err := json.Unmarshal(body, &layouts); err != nil {
		return "Error parsing response from API", err
	}

	numLayouts := len(layouts)
	var msg string
	if numLayouts > cmd_core.MaxLayouts {
		trimmed := strings.Join(layouts[:cmd_core.MaxLayouts], "\n")
		msg = fmt.Sprintf(
			"Found %d layouts: showing first %d\n"+
				"%s",
			numLayouts, cmd_core.MaxLayouts, trimmed,
		)
	} else {
		msg = fmt.Sprintf(
			"Found %d layouts:\n"+
				"%s",
			numLayouts, strings.Join(layouts, "\n"),
		)
	}
	return msg, nil
}

func (cmd List) Use() string {
	return "!akl list"
}

func (cmd List) Desc() string {
	return "list layout names"
}

package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/akl-infra/bot/internal/api"
)

type Layouts []string

const MaxLayouts = 20

func List(command string, args []string) (string, error) {
	res, err := api.Get("layouts")
	if err != nil {
		return "Error contacting API", err
	}
	defer res.Body.Close()
	var body []byte
	var layouts Layouts
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return "Error reading response from API", err
	}
	if err := json.Unmarshal(body, &layouts); err != nil {
		return "Error parsing response from API", err
	}

	numLayouts := len(layouts)
	var msg string
	if numLayouts > MaxLayouts {
		trimmed := strings.Join(layouts[:MaxLayouts], "\n")
		msg = fmt.Sprintf(
			"Found %d layouts: showing first %d\n"+
				"%s",
			numLayouts, MaxLayouts, trimmed,
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

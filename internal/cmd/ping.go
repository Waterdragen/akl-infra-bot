package cmd

import (
	"fmt"
	"strings"
)

func Ping(command string, args []string) (string, error) {
	return fmt.Sprintf("%s: [%s]", command, strings.Join(args, ",")), nil
}

package parser

import (
	"regexp"
	"strings"
)

func ParseQuotedArgs(str string) []string {
	re := regexp.MustCompile(`[^\s"']+|"([^"]*)"|'([^']*)'`)
	matches := re.FindAllString(str, -1)
	return matches
}

func ParseArgs(str string) []string {
	return strings.Fields(str)
}

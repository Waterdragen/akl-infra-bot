package parser

import (
	"fmt"
	"strings"
)

var prefixList = []string{"--", "—", "––"}

func ParseKwargs(arg string, argType KwargType, cmdKwargs map[string]KwargType) (map[string]Kwarg, error) {
	words := strings.Fields(arg)
	argIndex := 0
	for _, word := range words {
		isKwarg, err := isKwarg(cmdKwargs, word)
		if err != nil {
			return make(map[string]Kwarg), err
		}
		if isKwarg {
			break
		}
		argIndex++
	}

	// Make default map
	args := words[:argIndex]
	parsedKwargs := make(map[string]Kwarg)
	if argType == List {
		parsedKwargs["args"] = NewListValue(args)
	} else {
		parsedKwargs["args"] = NewStrValue(strings.Join(args, " "))
	}
	for kwName, kwType := range cmdKwargs {
		switch kwType {
		case Bool:
			parsedKwargs[kwName] = NewBoolValue(false)
		case List:
			parsedKwargs[kwName] = NewListValue([]string{})
		case Str:
			parsedKwargs[kwName] = NewStrValue("")
		}
	}

	words = words[argIndex:]
	lastInList := 0
	lastKwargType := List
	lastListKwarg := ""
	inList := false

	for index, word := range words {
		if ok, _ := isKwarg(cmdKwargs, word); !ok {
			continue
		}
		word = removeKwPrefix(word)
		kwType := cmdKwargs[word]

		// Encountered next keyword, stops previous list
		if inList {
			var value Kwarg
			if lastKwargType == List {
				value = NewListValue(words[lastInList:index])
			} else {
				value = NewStrValue(strings.Join(words[lastInList:index], " "))
			}
			parsedKwargs[lastListKwarg] = value
		}
		inList = kwType == List || kwType == Str

		// List | Str: Starts a new list after kwarg
		if !inList {
			parsedKwargs[word] = NewBoolValue(true)
		} else {
			lastKwargType = kwType
			lastListKwarg = word
			lastInList = index + 1
		}
	}

	// Close the last list
	if inList {
		var value Kwarg
		if lastKwargType == List {
			value = NewListValue(words[lastInList:])
		} else {
			value = NewStrValue(strings.Join(words[lastInList:], " "))
		}
		parsedKwargs[lastListKwarg] = value
	}

	return parsedKwargs, nil
}

func startsWithKwPrefix(word string) bool {
	for _, prefix := range prefixList {
		if strings.HasPrefix(word, prefix) {
			return true
		}
	}
	return false
}

func removeKwPrefix(word string) string {
	word = strings.ToLower(word)
	for _, prefix := range prefixList {
		if strings.HasPrefix(word, prefix) {
			return strings.TrimPrefix(word, prefix)
		}
	}
	return word
}

func isKwarg(kwargs map[string]KwargType, word string) (bool, error) {
	if !startsWithKwPrefix(word) {
		return false, nil
	}
	word = removeKwPrefix(word)
	_, found := kwargs[word]
	if found {
		return true, nil
	} else {
		return false, fmt.Errorf("invalid kwarg: `%s`", word)
	}
}

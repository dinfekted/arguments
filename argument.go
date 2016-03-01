package arguments

import (
	"fmt"
	"strconv"
	"strings"
)

type ArgumentType int

const (
	String  ArgumentType = iota
	Strings              = iota
	Integer              = iota
	Float                = iota
	Flag                 = iota
	Tail                 = iota
)

type Argument struct {
	Title       string
	Description string
	Type        ArgumentType
	Shortcut    string
	Required    bool
}

func (argument Argument) parseValue(values []string, pointer int) (
	interface{},
	int,
	error,
) {
	if argument.Type == String {
		return values[pointer+1], 1, nil
	}

	if argument.Type == Strings {
		shift := 0
		result := []string{}
		for {
			currentPointer := pointer + shift + 1
			if currentPointer >= len(values) {
				break
			}

			currentString := values[currentPointer]
			if strings.HasPrefix(currentString, "--") {
				break
			}
			result = append(result, currentString)
			shift += 1
		}

		return result, shift, nil
	}

	if argument.Type == Integer {
		result, err := strconv.Atoi(values[pointer+1])
		if err != nil {
			return nil, 0, err
		}

		return result, 1, nil
	}

	if argument.Type == Float {
		result, err := strconv.ParseFloat(values[pointer+1], 64)
		if err != nil {
			return nil, 0, err
		}

		return result, 1, nil
	}

	if argument.Type == Flag {
		if strings.HasPrefix(values[pointer], "--no-") {
			return false, 0, nil
		}

		return true, 0, nil
	}

	return nil, 0, fmt.Errorf("unknown argument type %s", argument.Type)
}

package arguments

import (
	"fmt"
	"strings"
)

type Arguments map[string]Argument

func (arguments Arguments) Parse(values []string) (Values, error) {
	pointer := 0
	result := Values{}
	for pointer < len(values) {
		parsed, newPointer, err := arguments.parseByName(values, result,
			pointer)

		if err != nil {
			return nil, err
		}

		if parsed {
			pointer = newPointer
			continue
		}

		parsed, newPointer, err = arguments.parseByShortcut(values, result,
			pointer)

		if err != nil {
			return nil, err
		}

		if parsed {
			pointer = newPointer
			continue
		}

		argument, key := arguments.findByType(Tail)
		if argument != nil {
			result[key] = values[pointer:]
			break
		}

		return nil, fmt.Errorf("unknown argument %s", values[pointer])
	}

	err := arguments.checkRequired(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (arguments Arguments) parseByName(
	values []string,
	result map[string]interface{},
	pointer int,
) (bool, int, error) {
	current := values[pointer]
	if !strings.HasPrefix(current, "--") {
		return false, 0, nil
	}

	name := current[2:]
	argument, key := arguments.findByKey(name)
	if argument == nil {
		if strings.HasPrefix(current, "--no-") {
			argument, key = arguments.findByKey(name[3:])
		}

		if argument == nil {
			return false, 0, fmt.Errorf("unknown option %s", current)
		}
	}

	value, shift, err := argument.parseValue(values, pointer)
	if err != nil {
		return false, 0, err
	}

	result[key] = value
	pointer += 1 + shift
	return true, pointer, nil
}

func (arguments Arguments) parseByShortcut(
	values []string,
	result map[string]interface{},
	pointer int,
) (bool, int, error) {
	current := values[pointer]

	if !strings.HasPrefix(current, "-") {
		return false, 0, nil
	}

	if len(current) == 2 {
		argument, key := arguments.findByShortcut(string(current[1:]))

		if argument == nil {
			return false, 0, fmt.Errorf("unknown flag %s", string(current))
		}

		value, shift, err := argument.parseValue(values, pointer)
		if err != nil {
			return false, 0, err
		}

		result[key] = value
		pointer += 1 + shift
		return true, pointer, nil
	}

	for _, value := range current[1:] {
		argument, key := arguments.findByShortcut(string(value))
		if argument == nil {
			return false, 0, fmt.Errorf("unknown flag -%s", string(value))
		}

		if argument.Type != Flag {
			return false, 0, fmt.Errorf("Can not mix flags with options")
		}

		result[key] = true
	}

	return true, pointer + 1, nil
}

func (arguments Arguments) findByKey(argumentName string) (
	*Argument,
	string,
) {
	for key, argument := range arguments {
		if key != argumentName {
			continue
		}

		return &argument, key
	}

	return nil, ""
}

func (arguments Arguments) findByType(argumentType ArgumentType) (
	*Argument,
	string,
) {
	for key, argument := range arguments {
		if argument.Type != argumentType {
			continue
		}

		return &argument, key
	}

	return nil, ""
}

func (arguments Arguments) findByShortcut(argumentShortcut string) (
	*Argument,
	string,
) {
	for key, argument := range arguments {
		if argument.Shortcut != argumentShortcut {
			continue
		}

		return &argument, key
	}

	return nil, ""
}

func (arguments Arguments) checkRequired(values map[string]interface{}) error {
	for key, argument := range arguments {
		if !argument.Required {
			continue
		}

		_, ok := values[key]
		if !ok {
			return fmt.Errorf("required option --%s not set", key)
		}
	}

	return nil
}

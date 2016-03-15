package arguments

import "errors"

var (
	types = map[string]ArgumentType{
		"string":  String,
		"strings": Strings,
		"integer": Integer,
		"float":   Float,
		"flag":    Flag,
	}
)

func Create(value interface{}) (Arguments, error) {
	mapping, ok := value.(map[interface{}]interface{})
	if !ok {
		return Arguments{}, errors.New("arguments value must be map")
	}

	result := Arguments{}
	for key, value := range mapping {
		keyString, ok := key.(string)
		if !ok {
			return Arguments{}, errors.New("key must be string")
		}

		arg, err := createArgument(value)
		if err != nil {
			return Arguments{}, err
		}

		result[keyString] = arg
	}

	return result, nil
}

func createArgument(value interface{}) (Argument, error) {
	mapping, ok := value.(map[interface{}]interface{})
	if !ok {
		return Argument{}, errors.New("argument value must be map")
	}

	result := Argument{}
	title, ok := mapping["title"]
	if ok {
		result.Title, ok = title.(string)
		if !ok {
			return Argument{}, errors.New("title must be string")
		}
	}

	description, ok := mapping["description"]
	if ok {
		result.Description, ok = description.(string)
		if !ok {
			return Argument{}, errors.New("description must be string")
		}
	}

	kind, ok := mapping["type"]
	if !ok {
		return Argument{}, errors.New("type must be set")
	}

	kindString, ok := kind.(string)
	if !ok {
		return Argument{}, errors.New("type must be string")
	}

	kindType, ok := types[kindString]
	if !ok {
		return Argument{}, errors.New("unknown argument type \"" + kindString + "\"")
	}

	result.Type = kindType

	shortcut, ok := mapping["shortcut"]
	if ok {
		result.Shortcut, ok = shortcut.(string)
		if !ok {
			return Argument{}, errors.New("shortcut must be string")
		}
	}

	required, ok := mapping["required"]
	if ok {
		result.Required, ok = required.(bool)
		if !ok {
			return Argument{}, errors.New("required must be bool")
		}
	}

	defaultValue, ok := mapping["default"]
	if ok {
		result.Default = defaultValue
		result.HasDefault = true
	}

	return result, nil
}

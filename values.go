package arguments

import "fmt"

type Values map[string]interface{}

func (values Values) Value(key string) (interface{}, bool) {
	result, ok := values[key]
	return result, ok
}

func (values Values) String(key, defaultValue string) (string, bool, error) {
	value, ok := values[key]
	if !ok {
		return defaultValue, false, nil
	}

	result, ok := value.(string)
	if !ok {
		return "", true, fmt.Errorf("argument %s should be string", key)
	}

	return result, true, nil
}

func (values Values) Strings(
	key string,
	defaultValue []string,
) ([]string, bool, error) {
	value, ok := values[key]
	if !ok {
		return defaultValue, false, nil
	}

	result, ok := value.([]string)
	if !ok {
		return []string{""}, false, fmt.Errorf("argument %s should be array "+
			"of strings", key)
	}

	return result, true, nil
}

func (values Values) Integer(
	key string,
	defaultValue int,
) (int, bool, error) {
	value, ok := values[key]
	if !ok {
		return defaultValue, false, nil
	}

	result, ok := value.(int)
	if !ok {
		return 0, true, fmt.Errorf("argument %s should be integer", key)
	}

	return result, true, nil
}

func (values Values) Float(
	key string,
	defaultValue float64,
) (float64, bool, error) {
	value, ok := values[key]
	if !ok {
		return defaultValue, false, nil
	}

	result, ok := value.(float64)
	if !ok {
		return 0, true, fmt.Errorf("argument %s should be float", key)
	}

	return result, true, nil
}

func (values Values) Boolean(
	key string,
	defaultValue bool,
) (bool, bool, error) {
	value, ok := values[key]
	if !ok {
		return defaultValue, false, nil
	}

	result, ok := value.(bool)
	if !ok {
		return false, true, fmt.Errorf("argument %s should be float", key)
	}

	return result, true, nil
}

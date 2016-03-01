package arguments

import "testing"

func TestParseNotReturnsError(test *testing.T) {
	arguments := Arguments{"key": {Type: String}}
	_, err := arguments.Parse([]string{"--key", "value"})
	if err != nil {
		test.Error(err)
	}
}

func TestParseReturnsString(test *testing.T) {
	arguments := Arguments{"key": {Type: String}}
	result, err1 := arguments.Parse([]string{"--key", "value"})
	item, found, err2 := result.String("key", "")
	if item != "value" || !found || err1 != nil || err2 != nil {
		test.Error(err1)
	}
}

func TestParseReturnsStringDefault(test *testing.T) {
	arguments := Arguments{"key": {Type: String}}
	result, err1 := arguments.Parse([]string{})
	item, found, err2 := result.String("key", "DEFAULT")
	if item != "DEFAULT" || found || err1 != nil || err2 != nil {
		test.Error(err1)
	}
}

func TestParseReturnsStrings(test *testing.T) {
	arguments := Arguments{"key": {Type: Strings}}
	result, err1 := arguments.Parse([]string{"--key", "value1", "value2"})
	item, found, err2 := result.Strings("key", nil)
	failed := item[0] != "value1" || item[1] != "value2"
	if failed || err1 != nil || err2 != nil || !found {
		test.Error(err1)
	}
}

func TestParseReturnsInteger(test *testing.T) {
	arguments := Arguments{"key": {Type: Integer}}
	result, err1 := arguments.Parse([]string{"--key", "100"})
	item, found, err2 := result.Integer("key", 0)
	if item != 100 || !found || err1 != nil || err2 != nil {
		test.Error(err1)
	}
}

func TestParseReturnsFloat(test *testing.T) {
	arguments := Arguments{"key": {Type: Float}}
	result, err1 := arguments.Parse([]string{"--key", "100.0"})
	item, found, err2 := result.Float("key", 0)
	if item != 100.0 || !found || err1 != nil || err2 != nil {
		test.Error(err1)
	}
}

func TestParseReturnsTrue(test *testing.T) {
	arguments := Arguments{"key": {Type: Flag}}
	result, err1 := arguments.Parse([]string{"--key"})
	item, found, err2 := result.Boolean("key", false)
	if item != true || !found || err1 != nil || err2 != nil {
		test.Error(err1)
	}
}

func TestParseReturnsFalse(test *testing.T) {
	arguments := Arguments{"key": {Type: Flag}}
	result, err1 := arguments.Parse([]string{"--no-key"})
	item, found, err2 := result.Boolean("key", true)
	if item != false || !found || err1 != nil || err2 != nil {
		test.Error(err1)
	}
}

func TestParseReturnsErrorOnUnknownOption(test *testing.T) {
	_, err := (Arguments{}).Parse([]string{"--key", "value"})
	if err == nil {
		test.Error(err)
	}
}

func TestParseReturnsErrorIfRequiredArgumentNotSet(test *testing.T) {
	arguments := Arguments{"key": {Type: String, Required: true}}
	_, err := arguments.Parse([]string{})
	if err == nil {
		test.Error(err)
	}
}

func TestParseReturnsStringByShortcut(test *testing.T) {
	arguments := Arguments{"key": {Shortcut: "k", Type: String}}
	result, err := arguments.Parse([]string{"-k", "value"})
	if result["key"] != "value" || err != nil {
		test.Error(err)
	}
}

func TestParseReturnsBooleansByShortcut(test *testing.T) {
	arguments := Arguments{
		"flag": {Shortcut: "f", Type: Flag},
		"bool": {Shortcut: "b", Type: Flag},
	}
	result, err := arguments.Parse(
		[]string{"-fb"},
	)
	if result["flag"] != true || result["bool"] != true || err != nil {
		test.Error(err)
	}
}

func TestParseReturnsErrorOnMixingBooleanAndString(test *testing.T) {
	arguments := Arguments{
		"key":   {Shortcut: "k", Type: Flag},
		"value": {Shortcut: "v", Type: String},
	}
	_, err := arguments.Parse([]string{"-kv", "value"})
	if err == nil {
		test.Fail()
	}
}

func TestParseReturnsErrorOnUnknwonShortcut(test *testing.T) {
	arguments := Arguments{}
	_, err := arguments.Parse([]string{"-k", "value"})
	if err == nil {
		test.Fail()
	}
}

func TestParseReturnsTail(test *testing.T) {
	arguments := Arguments{"tail": {Type: Tail}}
	result, err1 := arguments.Parse([]string{"value1", "value2"})
	tail, found, err2 := result.Strings("tail", nil)
	failed := err1 != nil || err2 != nil
	if tail[0] != "value1" || tail[1] != "value2" || !found || failed {
		test.Error(err1)
	}
}

func TestParseReturnsStringsAndString(test *testing.T) {
	arguments := Arguments{"strs": {Type: Strings}, "str": {Type: String}}
	result, err := arguments.Parse([]string{"--strs", "value1", "value2",
		"--str", "value"})
	strings := result["strs"].([]string)
	failed := strings[0] != "value1" || strings[1] != "value2"
	if failed || result["str"] != "value" || err != nil {
		test.Error(err)
	}
}

func TestParseReturnsIntegerAndString(test *testing.T) {
	arguments := Arguments{"int": {Type: Integer}, "str": {Type: String}}
	result, err := arguments.Parse([]string{"--int", "100", "--str", "value"})
	if result["int"] != 100 || result["str"] != "value" || err != nil {
		test.Error(err)
	}
}

func TestParseReturnsFloatAndString(test *testing.T) {
	arguments := Arguments{"float": {Type: Float}, "string": {Type: String}}
	result, err := arguments.Parse([]string{"--float", "100", "--string",
		"value"})
	if result["float"] != 100.0 || result["string"] != "value" || err != nil {
		test.Error(err)
	}
}

func TestParseReturnsTrueAndString(test *testing.T) {
	arguments := Arguments{"flag": {Type: Flag}, "string": {Type: String}}
	result, err := arguments.Parse([]string{"--flag", "--string", "value"})
	if result["flag"] != true || result["string"] != "value" || err != nil {
		test.Error(err)
	}
}

func TestParseReturnsFalseAndString(test *testing.T) {
	arguments := Arguments{"flag": {Type: Flag}, "string": {Type: String}}
	result, err := arguments.Parse([]string{"--no-flag", "--string", "value"})
	if result["flag"] != false || result["string"] != "value" || err != nil {
		test.Error(err)
	}
}

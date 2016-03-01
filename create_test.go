package arguments

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTestArguments(
	argument map[interface{}]interface{},
) (Arguments, error) {
	return Create(map[interface{}]interface{}{"argument": argument})
}

func TestCreateCreatesTitle(test *testing.T) {
	arguments, err := createTestArguments(map[interface{}]interface{}{
		"title": "TITLE",
		"type":  "string",
	})

	assert.NoError(test, err)
	assert.Equal(test, "TITLE", arguments["argument"].Title)
}

func TestCreateCreatesDescription(test *testing.T) {
	arguments, err := createTestArguments(map[interface{}]interface{}{
		"description": "DESCRIPTION",
		"type":        "string",
	})

	assert.NoError(test, err)
	assert.Equal(test, "DESCRIPTION", arguments["argument"].Description)
}

func TestCreateReturnsErrorIfNoTypeProvided(test *testing.T) {
	_, err := createTestArguments(map[interface{}]interface{}{})
	assert.Error(test, err)
}

func TestCreateReturnsErrorIfTypeIsUnknown(test *testing.T) {
	_, err := createTestArguments(map[interface{}]interface{}{"type": "WRONG"})
	assert.Error(test, err)
}

func TestCreateReturnsStringType(test *testing.T) {
	arguments, err := createTestArguments(map[interface{}]interface{}{
		"type": "string",
	})

	assert.NoError(test, err)
	assert.True(test, String == arguments["argument"].Type)
}

func TestCreateReturnsStringsType(test *testing.T) {
	arguments, err := createTestArguments(map[interface{}]interface{}{
		"type": "strings",
	})

	assert.NoError(test, err)
	assert.True(test, Strings == arguments["argument"].Type)
}

func TestCreateReturnsIntegerType(test *testing.T) {
	arguments, err := createTestArguments(map[interface{}]interface{}{
		"type": "integer",
	})

	assert.NoError(test, err)
	assert.True(test, Integer == arguments["argument"].Type)
}

func TestCreateReturnsFloatType(test *testing.T) {
	arguments, err := createTestArguments(map[interface{}]interface{}{
		"type": "float",
	})

	assert.NoError(test, err)
	assert.True(test, Float == arguments["argument"].Type)
}

func TestCreateReturnsFlagType(test *testing.T) {
	arguments, err := createTestArguments(map[interface{}]interface{}{
		"type": "flag",
	})

	assert.NoError(test, err)
	assert.True(test, Flag == arguments["argument"].Type)
}

func TestCreateCreatesShortcut(test *testing.T) {
	arguments, err := createTestArguments(map[interface{}]interface{}{
		"shortcut": "a",
		"type":     "string",
	})

	assert.NoError(test, err)
	assert.Equal(test, "a", arguments["argument"].Shortcut)
}

func TestCreateCreatesRequired(test *testing.T) {
	arguments, err := createTestArguments(map[interface{}]interface{}{
		"required": true,
		"type":     "string",
	})

	assert.NoError(test, err)
	assert.True(test, arguments["argument"].Required)
}

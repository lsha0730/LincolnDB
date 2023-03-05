package util

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	name     string
	args     interface{}
	expected interface{}
}

func TestValidateOp(t *testing.T) {
	tests := []Test{
		{
			name: "should pass normal operation",
			args: map[string]interface{}{
				"op": "read",
			},
			expected: nil,
		},
		{
			name:     "should fail missing value",
			args:     map[string]interface{}{},
			expected: errors.New("ERROR: Query is missing op property"),
		},
		{
			name: "should fail nonstring value",
			args: map[string]interface{}{
				"op": -24,
			},
			expected: errors.New("ERROR: Operation is not a string"),
		},
		{
			name: "should fail backspace string",
			args: map[string]interface{}{
				"op": "",
			},
			expected: errors.New("ERROR: Invalid operand clause"),
		},
		{
			name: "should fail nonquery string",
			args: map[string]interface{}{
				"op": "insert",
			},
			expected: errors.New("ERROR: Invalid operand clause"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := ValidateOp(test.args.(map[string]interface{}))
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestValidatePath(t *testing.T) {
	tests := []Test{
		{
			name: "should pass normal string",
			args: map[string]interface{}{
				"path": "`1234567890-=qwertyuiop[]asdfghjkl;'zxcvbnm,./?><:~!@#$%^&*()_+",
			},
			expected: nil,
		},
		{
			name: "should pass empty string",
			args: map[string]interface{}{
				"path": "",
			},
			expected: nil,
		},
		{
			name:     "should fail missing value",
			args:     map[string]interface{}{},
			expected: errors.New("ERROR: Query is missing path property"),
		},
		{
			name: "should fail nonstring value",
			args: map[string]interface{}{
				"path": -24,
			},
			expected: errors.New("ERROR: Path is not a string"),
		},
		{
			name: "should fail backspace string",
			args: map[string]interface{}{
				"path": "",
			},
			expected: errors.New("ERROR: Invalid path string"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := ValidatePath(test.args.(map[string]interface{}))
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestIsNormal(t *testing.T) {
	tests := []Test{
		{
			name:     "should pass normal string",
			args:     "`1234567890-=qwertyuiop[]asdfghjkl;'zxcvbnm,./?><:~!@#$%^&*()_+",
			expected: true,
		},
		{
			name:     "should pass empty string",
			args:     "",
			expected: true,
		},
		{
			name:     "should fail backspace string",
			args:     "",
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := IsNormal(test.args.(string))
			assert.Equal(t, test.expected, actual)
		})
	}
}

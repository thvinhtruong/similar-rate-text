package core

import (
	"reflect"
	"testing"
)

func TestVectorizeText(t *testing.T) {
	tests := []struct {
		text     string
		expected map[string]int
	}{
		{
			text:     "Hello world",
			expected: map[string]int{"hello": 1, "world": 1},
		},
		{
			text:     "Hello hello world",
			expected: map[string]int{"hello": 2, "world": 1},
		},
		{
			text:     "This is a test",
			expected: map[string]int{"this": 1, "is": 1, "a": 1, "test": 1},
		},
		{
			text:     "This is a test test",
			expected: map[string]int{"this": 1, "is": 1, "a": 1, "test": 2},
		},
	}

	for _, test := range tests {
		result := VectorizeText(test.text)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("VectorizeText(%s) = %v, expected %v", test.text, result, test.expected)
		}
	}
}

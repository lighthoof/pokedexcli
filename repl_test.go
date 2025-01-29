package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " CaMeL cAsE   dEtEcTeD",
			expected: []string{"camel", "case", "detected"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Result length differs from expected: %v vs %v", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Words do not match:")
				t.Errorf("Expected: %v", expectedWord)
				t.Errorf("Actual: %v", word)
			}
		}
	}
}

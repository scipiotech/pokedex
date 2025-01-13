package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Pikachu CHARGES forward",
			expected: []string{"pikachu", "charges", "forward"},
		},
	}
	//loop to check each test case
	for _, c := range cases {
		actual := cleanInput(c.input)

		// Check if lengths match
		if len(actual) != len(c.expected) {
			t.Errorf("Length mismatch: got %v, expected %v", len(actual), len(c.expected))
			continue
		}

		// Check each word
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("Word mismatch at position %d: got %v, expected %v", i, actual[i], c.expected[i])
			}
		}

	}
}

package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Salut je suis henry",
			expected: []string{"salut", "je", "suis", "henry"},
		},
		{
			input:    "Les amis ne sont Pas gentil",
			expected: []string{"les", "amis", "ne", "sont", "pas", "gentil"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice
		// if they don't match, use t.Errorf and continue to the next case
		if len(actual) != len(c.expected) {
			t.Errorf("actual of size %v and expected of size %v", len(actual), len(c.expected)) // error and continue here
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("expected %s and got %s", c.expected, actual)
			}
		}
	}
}

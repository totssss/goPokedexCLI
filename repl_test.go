package main

import "testing"

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
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q) returned %v; want %v", c.input, actual, c.expected)
			continue
		}

		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("cleanInput(%q) returned %v; want %v", c.input, actual, c.expected)
				break
			}
		}
	}
}

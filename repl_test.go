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
			input:    "   ",
			expected: []string{""},
		},
		{
			input:    "HELLO WORLD HI",
			expected: []string{"hello", "world", "hi"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("expected len %d, actual %d", len(c.expected), len(actual))
		}

		for i := range actual {
			word := actual[i]
			expected := c.expected[i]
			if word != expected {
				t.Errorf("got %q, expected %q", word, expected)
			}
		}
	}

}

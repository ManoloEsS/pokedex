package cli

import (
	"testing"
)

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
			input:    "hello   world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  hello  \n  world",
			expected: []string{"hello", "world"},
		},
		{
			input:    " \n hello  \n  world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  hello  \n  world  \n",
			expected: []string{"hello", "world"},
		},
		{
			input:    " Hello WorlD ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("%#v and %#v are not the same length", actual, c.expected)
			t.Fail()
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("%#v and %#v are not the same word", word, expectedWord)
				t.Fail()
			}
		}
	}
}

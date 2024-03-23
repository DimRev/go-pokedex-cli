package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "hello world",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "HELLO world",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)
		fmt.Println("--------------------------")
		fmt.Printf("--Input--\n%v\n--Expected--\n%v\n--Actual--\n%v\n", cs.input, cs.expected, actual)
		if len(actual) != len(cs.expected) {
			t.Errorf("The lengths are not equal: %v vs %v ",
				len(actual),
				len(cs.expected))
			fmt.Println("\n\tFAIL")
			continue
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.expected[i]
			if actualWord != expectedWord {
				t.Errorf("%v does not equal %v", actualWord, expectedWord)
				fmt.Println("\n\tFAIL")
				continue
			}
		}
		fmt.Println("\n\tPASS")
		fmt.Println("--------------------------")
	}
}

package main

import "testing"
import "fmt"

func TestLookAndSay(t *testing.T) {

	input := "1"
	results := []string{"11", "21", "1211", "111221", "312211"}

	for _, r := range results {
		input = lookAndSay(input)
		fmt.Println("OK", input, r)
		if input != r {
			t.Errorf("NOK, got: %s, want: %s.", input, r)
		}
	}
}

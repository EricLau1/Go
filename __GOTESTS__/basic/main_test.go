package main

import "testing"

func TestSquare(t *testing.T) {
	if Square(2) != 4 {
		t.Error("Expected 2 ^ 2 to equal 4")
	}
}

func TestTableSquare(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{3, 9},
		{5, 25},
		{80, 6400},
	}

	for _, test := range tests {
		if output := Square(test.input); output != test.expected {
			t.Errorf("Test Failed: %d inputed, expected %d, received: %d", test.input, test.expected, output)
		}
	}
}

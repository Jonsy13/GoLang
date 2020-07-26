package main

import "testing"

func TestCalulate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 2 + 2 to equal to 4")
	}
}

func TestTable(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test Failed: {} inputtted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}

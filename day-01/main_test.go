package main

import (
	"testing"
)

func TestCalcFuelRequirement(t *testing.T) {
	var cases = []struct {
		input    int
		expected int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, tt := range cases {
		got := calcFuelRequirement(tt.input)
		if got != tt.expected {
			t.Errorf("calcFuelRequirement(%d) got %d, want %d", tt.input, got, tt.expected)
		}
	}
}

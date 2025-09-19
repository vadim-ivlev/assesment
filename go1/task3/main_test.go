package main

import "testing"

func TestCalculateAverage(t *testing.T) {

	tests := []struct {
		input    string
		expected string
	}{
		{"87 92 56 104 75 89", "79.8"},
		{"-1 19.1 18,2", "Нет корректных оценок"},
		{"80 -10 90 110", "85.0"},
		{"100 0 50", "50.0"},
		{"101 102 103", "Нет корректных оценок"},
		{"0 0 0", "0.0"},
		{"100 100 100", "100.0"},
		{"50 50 50 50 50", "50.0"},
		{"1 2 3 4 5", "3.0"},
		{"-100 100 0", "50.0"},
	}

	for _, tt := range tests {
		result := calculateAverage(tt.input)
		if result != tt.expected {
			t.Errorf("input: %q, expected: %q, got: %q", tt.input, tt.expected, result)
		}
	}
}

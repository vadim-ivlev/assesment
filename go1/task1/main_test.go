package main

import "testing"

func TestGetHTTPStatusCategory(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{200, "Success"},
		{404, "Client Error"},
		{100, "Informational"},
		{150, "Informational"},
		{201, "Success"},
		{301, "Redirection"},
		{500, "Server Error"},
		{599, "Server Error"},
	}

	for _, test := range tests {
		result := GetHTTPStatusCategory(test.input)
		if result != test.expected {
			t.Errorf("For input %d, expected %q, got %q", test.input, test.expected, result)
		}
	}
}

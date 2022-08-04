package main

import "testing"

func TestIsEven(t *testing.T) {
	//これを初めに書くことで並列でテストを実行することができる
	t.Parallel()

	testCase := []struct {
		name     string
		input    int
		expected bool
	}{
		{"Positive Odd", 1, false},
		{"Positive Even", 2, true},
		{"Negative Odd", -1, false},
		{"Negative Even", -2, true},
	}

	//
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			got := IsEven(tt.input)
			if got != tt.expected {
				t.Errorf("expected: %v, got: %v\n", tt.expected, got)
			}
		})
	}
}

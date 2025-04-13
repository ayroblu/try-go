package main

import (
	"fmt"
	"testing"
)

func TestExample(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{2, 3, 5},
		{0, 0, 0},
		{-1, 1, 0},
		{-5, -3, -8},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v + %v == %v", tt.a, tt.b, tt.expected), func(t *testing.T) {
			if tt.a+tt.b != tt.expected {
				t.Error("Invalid run", tt.a, tt.b, "!=", tt.expected)
			}
		})
	}
	t.Error("invalid")
}

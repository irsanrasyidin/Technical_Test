package main

import (
	"testing"
	"errors"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{1, 2, 3},
		{-1, 1, 0},
		{0, 0, 0},
	}

	for _, tt := range tests {
		t.Run("Add", func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		a, b          int
		expected      int
		expectedError error
	}{
		{6, 2, 3, nil},
		{5, 0, 0, errors.New("division by zero")},
		{10, -2, -5, nil},
	}

	for _, tt := range tests {
		t.Run("Divide", func(t *testing.T) {
			result, err := Divide(tt.a, tt.b)
			if err != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("Divide(%d, %d) error = %v; want %v", tt.a, tt.b, err, tt.expectedError)
			}
			if result != tt.expected {
				t.Errorf("Divide(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

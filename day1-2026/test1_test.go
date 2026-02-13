package main

import "testing"

func TestAdd(t *testing.T) {
	result := add(1, 2)
	expected := 3
	if result != expected {
		t.Errorf("Expected 3, got %d", result)
	}
}
func TestDivide(t *testing.T) {
	// Test normal division
	result, err := divide(10, 2)
	if err != nil {
		t.Errorf("divide(10, 2) returned error: %v", err)
	}
	if result != 5 {
		t.Errorf("divide(10, 2) = %d; want 5", result)
	}

	// Test division by zero
	_, err = divide(10, 0)
	if err == nil {
		t.Error("divide(10, 0) should return error, got nil")
	}
}

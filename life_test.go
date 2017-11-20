package main

import "testing"

// Test NewUniverse function that uses make to allocate and return a Universe
// with Height rows and Width columns per row.
func TestNewUniverse(t *testing.T) {
	u := NewUniverse()
	if len(u) != y {
		t.Errorf("Newly initialized universe is %v rows high. Want %v rows.", len(u), y)
	}

	for row := range u {
		if len(u[row]) != x {
			t.Errorf("Newly initialized universe is %v columns wide. Want %v columns.", len(u[0]), x)
		}
	}
}

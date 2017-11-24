package main

import (
	"math/rand"
	"testing"
)

var fmtRowTests = []struct {
	row  []bool
	want string
}{
	{[]bool{false, true, false}, " * \n"},
	{[]bool{true, false, true}, "* *\n"},
}

// Test NewUniverse function that uses make to allocate and return a Universe
// with Height rows and Width columns per row.
func TestNewUniverse(t *testing.T) {
	t.Parallel()
	x, y := 80, 15

	u := NewUniverse(x, y)
	if len(u) != y {
		t.Errorf("Newly initialized universe is %v rows high. Want %v rows.", len(u), y)
	}

	for row := range u {
		if len(u[row]) != x {
			t.Errorf("Newly initialized universe is %v columns wide. Want %v columns.", len(u[0]), x)
		}
	}
}

func TestFormatRow(t *testing.T) {
	t.Parallel()

	for _, tt := range fmtRowTests {
		got := formatRow(tt.row)
		if got != tt.want {
			t.Errorf("formatRow(%v) => %v; want %v", tt.row, got, tt.want)
		}
	}
}

func TestUniverseShow(t *testing.T) {
	t.Parallel()
	x, y := 80, 15

	u := NewUniverse(x, y)
	got := len(u.Show())
	want := (x * y) + y

	if got != want {
		t.Errorf("NewUniverse is a string %v characters long; want %v characters long.", got, want)
	}
}

func TestUniverseSeed(t *testing.T) {
	t.Parallel()
	x, y := 80, 15

	rand.Seed(1)
	percentLive := 0.25
	totalCells := x * y
	var liveActual int
	liveExpected := 329 // this value depends on rand.Seed(1), totalCells = 1200, and % live of 0.25

	u := NewUniverse(x, y)
	u.Seed(percentLive)

	for _, row := range u {
		for _, cell := range row {
			if cell {
				liveActual++
			}
		}
	}

	if liveActual != liveExpected {
		t.Errorf("After Seeding, new universe has %v/%v cells alive; want %v/%v.",
			liveActual, totalCells, liveExpected, totalCells)
	}
}

func TestAlive(t *testing.T) {
	t.Parallel()
	// x, y := 5, 5

	// u := NewUniverse(x, y)

}

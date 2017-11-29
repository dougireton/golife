package main

import (
	"fmt"
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

// aliveTests assume a 4x4 universe
var aliveTests = []struct {
	x    int
	y    int
	want bool
}{
	// cell a
	{0, -1, true},   // N
	{1, -1, false},  // NE
	{1, 0, true},    // E
	{1, 1, false},   // SE
	{0, 1, true},    // S
	{-1, 1, false},  // SW
	{-1, 0, true},   // W
	{-1, -1, false}, // NW

	// cell b
	{3, 2, true},  // N
	{4, 2, false}, // NE
	{4, 3, true},  // E
	{4, 4, false}, // SE
	{3, 4, true},  // S
	{2, 4, false}, // SW
	{2, 3, true},  // W
	{2, 2, false}, // NW
}

// Test NewUniverse function which uses make to allocate and return a Universe
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

// TestFormatRow verifies the FormatRow helper function formats a Universe
// row for printing to the screen
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

	rand.Seed(1)

	x, y := 80, 15
	total := x * y
	percentLive := 0.25
	var got int
	want := 329 // this value depends on rand.Seed(1), total = 1200, and % live of 0.25

	u := NewUniverse(x, y)
	u.Seed(percentLive)

	for _, row := range u {
		for _, cell := range row {
			if cell {
				got++
			}
		}
	}

	if got != want {
		t.Errorf("After Seeding, new universe has %v/%v cells alive; want %v/%v.",
			got, total, want, total)
	}
}

// TestAlive verifies the Universe.Alive method.
func TestAlive(t *testing.T) {
	t.Parallel()
	x, y := 4, 4

	u := NewUniverse(x, y)
	fmt.Println("u len:", len(u[0]))

	// populate the universe
	// Cells are initialized to false. Set specified ones to alive/true.
	u[0][1] = true // E
	u[0][3] = true // S
	u[1][0] = true // S
	u[2][3] = true // N
	u[3][0] = true // E
	u[3][2] = true // W

	for _, tt := range aliveTests {
		got := u.Alive(tt.x, tt.y)

		if got != tt.want {
			t.Errorf("u.Alive(%v, %v) => %v; want %v", tt.x, tt.y, got, tt.want)
		}
	}

}

func TestNeighbors(t *testing.T) {

}

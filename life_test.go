package main

import (
	"fmt"
	"math/rand"
	"testing"
)

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
	fmtRowTests := []struct {
		row  []bool
		want string
	}{
		{[]bool{false, true, false}, " * \n"},
		{[]bool{true, false, true}, "* *\n"},
	}

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

	var got int

	rand.Seed(1)

	x, y := 80, 15
	total := x * y
	percentLive := 0.25
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

	// aliveTests assume a 4x4 universe
	aliveTests := []struct {
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

	u := NewUniverse(4, 4)

	// populate the universe
	// Cells are initialized to false. Set specified ones to alive/true.
	//   +---------------+
	// 0 | F | T | F | T |
	// 1 | T | F | T | F |
	// 2 | F | T | F | T |
	// 3 | T | F | T | F |
	//   +---------------+
	//     0   1   2   3

	//y, x
	u[0][1] = true
	u[0][3] = true
	u[1][0] = true
	u[1][2] = true
	u[2][1] = true
	u[2][3] = true
	u[3][0] = true
	u[3][2] = true

	for _, tc := range aliveTests {
		got := u.Alive(tc.x, tc.y)

		if got != tc.want {
			t.Errorf("u.Alive(%v, %v) => %v; want %v", tc.x, tc.y, got, tc.want)
		}
	}

}

func TestNeighbors(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		x    int
		y    int
		want int
	}{
		{"a", 0, 0, 3},
		{"b", 3, 1, 1},
	}

	u := NewUniverse(4, 4)

	// populate the universe
	// Cells are initialized to false. Set specified ones to alive/true.
	//   +---------------+
	// 0 | a | T | F | F |
	// 1 | T | F | F | b |
	// 2 | F | T | F | F |
	// 3 | T | F | F | F |
	//   +---------------+
	//     0   1   2   3
	// Cell "a" has 3 alive neighbors
	// Cell "b" has 1 alive neighbor

	//y, x
	u[0][1] = true
	u[1][0] = true
	u[2][1] = true
	u[3][0] = true

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Cell %v at position x=%v, y=%v", tc.name, tc.x, tc.y), func(t *testing.T) {
			if got := u.Neighbors(tc.x, tc.y); got != tc.want {
				t.Errorf("got %v; want %v", got, tc.want)
			}
		})
	}
}

// A live cell with fewer than two live neighbors dies
//   +-----------+
// 0 | F | F | F |
// 1 | F | a | F |
// 2 | F | F | F |
//   +-----------+
//     0   1   2
func TestNextLiveCellFewerThanTwoLiveNeighborsDies(t *testing.T) {
	t.Parallel()

	u := NewUniverse(3, 3)

	u[1][1] = true
	want := false
	if got := u.Next(1, 1); got != want {
		t.Errorf("u.Next(1, 1) => %v; want %v", got, want)
	}

}

// A live cell with two or three live neighbors lives on to the next generation.
//   +-----------+
// 0 | F | T | F |
// 1 | F | a | T |
// 2 | F | F | F |
//   +-----------+
//     0   1   2
func TestNextLiveCellTwoLiveNeighborsLives(t *testing.T) {
	t.Parallel()

	u := NewUniverse(3, 3)

	u[0][1] = true
	u[1][1] = true
	u[1][2] = true
	want := true
	if got := u.Next(1, 1); got != want {
		t.Errorf("u.Next(1, 1) => %v; want %v", got, want)
	}
}

// A live cell with two or three live neighbors lives on to the next generation.
//   +-----------+
// 0 | F | T | F |
// 1 | F | a | T |
// 2 | T | F | F |
//   +-----------+
//     0   1   2
func TestNextLiveCellThreeLiveNeighbors(t *testing.T) {
	t.Parallel()

	u := NewUniverse(3, 3)

	u[0][1] = true
	u[1][1] = true
	u[1][2] = true
	u[2][0] = true
	want := true
	if got := u.Next(1, 1); got != want {
		t.Errorf("u.Next(1, 1) => %v; want %v", got, want)
	}
}

// A live cell with more than three live neighbors dies.
//   +-----------+
// 0 | T | T | T |
// 1 | T | a | F |
// 2 | T | F | F |
//   +-----------+
//     0   1   2
func TestNextLiveCellMoreThanThreeLiveNeighborsDies(t *testing.T) {
	t.Parallel()

	u := NewUniverse(3, 3)

	u[0][0] = true
	u[0][1] = true
	u[0][2] = true
	u[1][1] = true
	u[1][0] = true
	u[2][0] = true
	want := false
	if got := u.Next(1, 1); got != want {
		t.Errorf("u.Next(1, 1) => %v; want %v", got, want)
	}
}

// A dead cell with exactly three live neighbors becomes a live cell.
//   +-----------+
// 0 | T | T | T |
// 1 | F | a | F |
// 2 | F | F | F |
//   +-----------+
//     0   1   2
func TestNextDeadCellExactlyThreeLiveNeighborsResurrects(t *testing.T) {
	t.Parallel()

	u := NewUniverse(3, 3)

	u[0][0] = true
	u[0][1] = true
	u[0][2] = true
	want := true
	if got := u.Next(1, 1); got != want {
		t.Errorf("u.Next(1, 1) => %v; want %v", got, want)
	}
}

// A dead cell with != 3 live neighbors stays dead.
//   +-----------+
// 0 | T | F | F |
// 1 | F | a | F |
// 2 | F | F | F |
//   +-----------+
//     0   1   2
func TestNextDeadCellWithoutThreeLiveNeighborsStaysDead(t *testing.T) {
	t.Parallel()

	u := NewUniverse(3, 3)

	u[0][0] = true
	want := false
	if got := u.Next(1, 1); got != want {
		t.Errorf("u.Next(1, 1) => %v; want %v", got, want)
	}
}

//    Current          Next
//   +-----------+    +-----------+
// 0 | T | T | F |  0 | T | T | F |
// 1 | T | F | F |  1 | T | T | F |
// 2 | F | F | T |  2 | F | F | T |
//   +-----------+    +-----------+
//     0   1   2        0   1   2
func TestStepUniverse(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		x    int
		y    int
		want bool
	}{
		{0, 0, true},
		{0, 1, true},
		{0, 2, false},
		{1, 0, true},
		{1, 1, false},
		{1, 2, false},
		{2, 0, false},
		{2, 1, false},
		{2, 2, true},
	}

	a := NewUniverse(3, 3)
	b := NewUniverse(3, 3)

	//y  x
	a[0][0] = true
	a[0][1] = true
	a[1][0] = true
	a[2][2] = true

	Step(a, b)

	for _, tc := range testCases {
		if got := b[tc.y][tc.x]; got != tc.want {
			t.Errorf("b[%v][%v] => %v; want %v", tc.y, tc.x, got, tc.want)
		}
	}

}

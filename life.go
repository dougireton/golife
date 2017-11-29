package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

type Universe [][]bool

func init() {
	rand.Seed(time.Now().UnixNano())
}

// NewUniverse creates a new, empty Universe of y columns high by x columns wide.
func NewUniverse(x, y int) Universe {
	u := make([][]bool, y)
	cells := make([]bool, x*y)
	for row := range u {
		u[row], cells = cells[:x:x], cells[x:]
	}
	return u
}

// formatRow formats a row in the Universe.
// Helper function for the Show function.
func formatRow(r []bool) string {
	var b bytes.Buffer
	for _, cell := range r {
		if cell {
			b.WriteString("*")
		} else {
			b.WriteString(" ")
		}
	}
	b.WriteString("\n")
	return b.String()
}

// Show creates a printable universe.
// Live cells are shown with an asterisk and dead cells
// are shown with a space.
func (u Universe) Show() string {
	var b bytes.Buffer
	for _, row := range u {
		b.WriteString(formatRow(row))
	}
	return b.String()
}

func initializeCell(p float64) bool {
	num := rand.Intn(100) + 1
	return num <= int(p*100)
}

// Seed randomly sets p % of the cells to alive (true).
// For a 1200 cell universe with p = 25%, the number of live cells should be roughly 300.
func (u Universe) Seed(p float64) {
	for y, row := range u {
		for x := range row {
			u[y][x] = initializeCell(p)
		}
	}
}

// Alive returns true if the cell at position x, y is alive, false if dead.
//
// The universe needs to wrap around, such that all cells, including edge
// and corner cells, have eight neighbors. For example, given a cell, "a", in the table below,
// its eight neighbors are shown by eight compass points (N, S, E, W, etc).
// Therefore, for cell "a" at position (x=0, y=0), its NW neighbor is at (x=3, y=3),
// and its W neighbor is at (x=3, y=0).
//   +-----------------+
// 0 | a | E  |   | W  |
// 1 | S | SE |   | SW |
// 2 |   |    |   |    |
// 3 | N | NE |   | NW |
//   + ----------------+
//     0   1    2   3
func (u Universe) Alive(x, y int) bool {
	height := len(u)
	width := len(u[0])

	// Use modulus to make the rows/colums "wrap around".
	// See here for a simplified example: https://play.golang.org/p/q68IKzNof1
	y = (y + height) % height
	x = (x + width) % width

	return u[y][x]
}

// Neighbors returns the number of live neighbors for a given cell, from zero to eight.
func (u Universe) Neighbors(x, y int) int {
	return 0
}

func main() {
	foo := NewUniverse(80, 15)
	fmt.Print(foo.Show())
}

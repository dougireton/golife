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
	for i, row := range u {
		for j := range row {
			u[i][j] = initializeCell(p)
		}
	}
}

// Alive returns true if the cell at position x, y is alive, false if dead.
func (u Universe) Alive(x, y int) bool {
	return true
}

func main() {
	foo := NewUniverse(80, 15)
	fmt.Print(foo.Show())
}

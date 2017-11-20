package main

import "fmt"

const (
	x = 80
	y = 15
)

type Universe [][]bool

func NewUniverse() Universe {
	u := make([][]bool, y)
	cells := make([]bool, x*y)
	for row := range u {
		u[row], cells = cells[:x:x], cells[x:]
	}
	return u
}

func main() {

	foo := NewUniverse()
	fmt.Println(foo)
}

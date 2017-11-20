package main

import "fmt"

const (
	height = 15
	width  = 80
)

type Universe [][]bool

func NewUniverse() Universe {
	u := make([][]bool, height)
	for row := 0; row < height; row++ {
		u[row] = make([]bool, 80)
	}
	return u
}

func main() {

	foo := NewUniverse()
	fmt.Println(foo)
}

# Game of Life

Chapter 20 capstone exercise from Nathan Youngman's "Get Programming with Go"

John Conway’s Game of Life is a simulation played out on a two dimensional grid of cells.
As such, this challenge focuses on slices.
In each generation, cells live or die based on their surrounding neighbors.
Each cell has eight neighbors, which are adjacent in the horizontal, vertical, and diagonal directions.

A live cell with less than two live neighbors dies.
A live cell with two or three live neighbors lives on to the next generation.
A live cell with more than three live neighbors dies.
A dead cell with exactly three live neighbors becomes a live cell.

Nathan Youngman. Get Programming with Go MEAP V11 (Kindle Locations 3812-3819). Manning Publications Co.. Kindle Edition.

## Unit Tests

Full test coverage is included.

```
go test
```
package main

import (
	"AOC/h"
	"fmt"
)

func main() {
	lines := h.GetLinesAsSlice()
	grid := h.ConvertLinesToGrid(lines)
	fmt.Printf("%c", grid[2][1])
	h.PrintGrid(grid)
}

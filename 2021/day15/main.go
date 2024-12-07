package main

import (
	"AOC/h"
	"fmt"
)

func main() {
	fmt.Println("Start")
	part1()
	part2()
}

func part1() {
	sum := 0
	lines := h.GetLinesAsSlice()
	grid := h.ConvertLinesToGrid(lines)

	startPos := h.Point{0, 0}
	endPos := h.Point{len(grid), len(grid[0])}

	fmt.Println(sum)
}

func part2() {

}

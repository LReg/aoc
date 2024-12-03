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
	lines := h.GetLinesAsSlice()
	grid := h.ConvertLinesToGrid(lines)
	height := len(grid[0])
	sum := 0
	x := 0
	for i := 0; i < height; i++ {
		if string(grid[x%len(grid)][i]) == "#" {
			sum++
		}
		x += 3
	}
	fmt.Println(sum)
}

func part2() {
	mult := 1
	lines := h.GetLinesAsSlice()
	grid := h.ConvertLinesToGrid(lines)
	height := len(grid[0])
	for _, data := range [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}} {
		sum := 0
		x := 0
		for i := 0; i < height; i += data[1] {
			if string(grid[x%len(grid)][i%len(grid[0])]) == "#" {
				sum++
			}
			x += data[0]
		}
		mult *= sum
	}
	fmt.Println(mult)
}

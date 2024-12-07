package main

import (
	"AOC/h"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Start")
	part1()
	part2()
}

func part1() {
	lines := h.GetLinesAsSlice()
	grid := h.ConvertLinesToGrid(lines)

	startPos := h.Point{0, 0}
	endPos := h.Point{len(grid) - 1, len(grid[0]) - 1}

	_, l := grid.DijkstraPosNum(startPos, endPos)

	fmt.Println(l)
}

func part2() {
	lines := h.GetLinesAsSlice()
	grid := h.ConvertLinesToGrid(lines)

	updatedGrid := make(h.Grid, len(grid)*5)
	for i := 0; i < len(grid)*5; i++ {
		updatedGrid[i] = make([]byte, len(grid[0])*5)
	}

	updatedGrid.ForEachPoint(func(p h.Point) {
		additionerX := p.X / len(grid)
		additionerY := p.Y / len(grid[0])
		gp := h.Point{
			p.X % len(grid),
			p.Y % len(grid[0]),
		}
		pnum := grid.AtNum(gp)
		addReset := 0
		if pnum+additionerX+additionerY > 9 {
			addReset = 1
		}
		updatedGrid.Set(p, strconv.Itoa((pnum+additionerY+additionerX)%10 + addReset)[0])
	})

	_, l := updatedGrid.DijkstraPosNum(h.Point{0, 0}, h.Point{len(updatedGrid) - 1, len(updatedGrid[0]) - 1})
	fmt.Println(l)
}

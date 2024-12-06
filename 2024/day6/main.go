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
	dir := h.NORTH
	pos := h.Point{0, 0}
	grid.ForEachPoint(func(p h.Point) {
		if grid.At(p) == '^' {
			pos = p
		}
	})
	for {
		if !h.IsPointInGrid(grid, pos) {
			break
		}
		if !h.IsPointInGrid(grid, pos.Relative(dir)) {
			sum++
			break
		}
		if grid.At(pos.Relative(dir)) == '#' {
			if dir == h.NORTH {
				dir = h.EAST
			} else if dir == h.EAST {
				dir = h.SOUTH
			} else if dir == h.SOUTH {
				dir = h.WEST
			} else if dir == h.WEST {
				dir = h.NORTH
			}
			continue
		}
		prev := pos
		pos = pos.Relative(dir)
		grid.Set(prev, 'X')
		if !h.IsPointInGrid(grid, pos) {
			break
		}
	}

	grid.ForEachPoint(func(p h.Point) {
		if grid.At(p) == 'X' {
			sum++
		}
	})
	fmt.Println(sum)
}

func part2() {
	sum := 0
	lines := h.GetLinesAsSlice()
	grid := h.ConvertLinesToGrid(lines)

	grid.ForEachPoint(func(p h.Point) {

		dir := h.NORTH
		pos := h.Point{0, 0}
		grid.ForEachPoint(func(p h.Point) {
			if grid.At(p) == '^' {
				pos = p
			}
		})

		ob := p
		v := make([]h.Point, 0)

		for i := 0; true; i++ {
			if !h.IsPointInGrid(grid, pos) {
				break
			}
			if !h.IsPointInGrid(grid, pos.Relative(dir)) {
				break
			}
			if i == len(grid)*len(grid[0]) {
				sum++
				break
			}

			if grid.At(pos.Relative(dir)) == '#' || pos.Relative(dir) == ob {
				if dir == h.NORTH {
					dir = h.EAST
				} else if dir == h.EAST {
					dir = h.SOUTH
				} else if dir == h.SOUTH {
					dir = h.WEST
				} else if dir == h.WEST {
					dir = h.NORTH
				}
				continue
			}
			v = append(v, pos)
			pos = pos.Relative(dir)
		}
	})
	fmt.Println(sum)
}

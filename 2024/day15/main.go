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
	H := 0
	dirs := []int{}
	pos := h.Point{}
	for _, l := range lines {
		if len(l) == 0 {
			break
		}
		H++
	}
	grid := h.ConvertLinesToGrid(lines[:H])
	for _, l := range lines[H+1:] {
		for _, r := range l {
			if r == '>' {
				dirs = append(dirs, h.EAST)
			} else if r == '<' {
				dirs = append(dirs, h.WEST)
			} else if r == '^' {
				dirs = append(dirs, h.NORTH)
			} else if r == 'v' {
				dirs = append(dirs, h.SOUTH)
			}
		}
	}
	grid.ForEachPoint(func(p h.Point) {
		if grid.At(p) == '@' {
			pos = p
			grid.Set(p, '.')
		}
	})

	for _, dir := range dirs {
		canMove := false
		h.WalkThroughLine(grid, dir, pos, func(p h.Point, i int) bool {
			if !p.IsInGrid(grid) {
				return false
			}
			if grid.At(p) == '#' {
				return false
			}
			if grid.At(p) == '.' && p != pos {
				canMove = true
				return false
			}
			return true
		})
		if !canMove {
			continue
		}
		pos = pos.Relative(dir)
		if grid.At(pos) == 'O' {
			pr := pos
			grid.Set(pr, '.')
			pr = pr.Relative(dir)
			for grid.At(pr) == 'O' {
				pr = pr.Relative(dir)
			}
			grid.Set(pr, 'O')
		}
	}
	grid.ForEachPoint(func(p h.Point) {
		if grid.At(p) == 'O' {
			sum += 100*p.Y + p.X
		}
	})
	fmt.Println(sum)
}

func part2() {

}

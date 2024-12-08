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
	sum := 0
	lines := h.GetLinesAsSlice()
	grid := h.ConvertLinesToGrid(lines)
	for step := 1; step <= 3; step++ {
		h.PrintGrid(grid)
		fmt.Println("--- STEP ---")

		flashed := make([]h.Point, 0)
		// inc + 1 of F
		grid.ForEachPoint(func(p h.Point) {
			prevNum := grid.AtNum(p)
			if prevNum < 9 {
				grid.Set(p, strconv.Itoa(prevNum + 1)[0])
			} else {
				grid.Set(p, 'F')
				flashed = append(flashed, p)
			}
		})

		h.PrintGrid(grid)
		println()

		for {
			flashedAgain := false
			fmt.Println("flashed ones: ", flashed)
			for _, fl := range flashed {
				fmt.Println("handling flashed one: ", fl)
				for _, n := range fl.Neighbours() {
					if n.IsInGrid(grid) {
						prev := grid.At(n)
						prevNum := grid.AtNum(n)
						if prev == 'F' {
							continue
						} else if prevNum == 9 {
							flashedAgain = true
							flashed = append(flashed, n)
							grid.Set(n, 'F')
						} else {
							grid.Set(n, strconv.Itoa(prevNum + 1)[0])
						}
					}
				}
				fmt.Println("flashed after delete", flashed)
			}
			h.PrintGrid(grid)
			println()
			if !flashedAgain {
				break
			}
		}

		grid.ForEachPoint(func(p h.Point) {
			if grid.At(p) == 'F' {
				grid.Set(p, '0')
			}
		})
	}
	fmt.Println(sum)
}

func part2() {

}

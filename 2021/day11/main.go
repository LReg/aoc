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
	for step := 1; step <= 100; step++ {
		flashed := make([]h.Point, 0)
		// inc + 1 of F
		grid.ForEachPoint(func(p h.Point) {
			prevNum := grid.AtNum(p)
			if prevNum < 9 {
				grid.Set(p, strconv.Itoa(prevNum + 1)[0])
			} else {
				grid.Set(p, 'F')
				sum++
				flashed = append(flashed, p)
			}
		})

		for {
			flashedAgain := make([]h.Point, 0)
			for _, fl := range flashed {
				for _, n := range fl.Neighbours() {
					if n.IsInGrid(grid) {
						prev := grid.At(n)
						prevNum := grid.AtNum(n)
						if prev == 'F' {
							continue
						} else if prevNum == 9 {
							flashedAgain = append(flashedAgain, n)
							sum++
							grid.Set(n, 'F')
						} else {
							grid.Set(n, strconv.Itoa(prevNum + 1)[0])
						}
					}
				}
			}

			if len(flashedAgain) == 0 {
				break
			}
			flashed = flashedAgain
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
	lines := h.GetLinesAsSlice()
	grid := h.ConvertLinesToGrid(lines)
	for step := 1; true; step++ {
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

		for {
			flashedAgain := make([]h.Point, 0)
			for _, fl := range flashed {
				for _, n := range fl.Neighbours() {
					if n.IsInGrid(grid) {
						prev := grid.At(n)
						prevNum := grid.AtNum(n)
						if prev == 'F' {
							continue
						} else if prevNum == 9 {
							flashedAgain = append(flashedAgain, n)
							grid.Set(n, 'F')
						} else {
							grid.Set(n, strconv.Itoa(prevNum + 1)[0])
						}
					}
				}
			}

			if len(flashedAgain) == 0 {
				break
			}
			flashed = flashedAgain
		}

		anyPNotFlash := false
		grid.ForEachPoint(func(p h.Point) {
			if grid.At(p) == 'F' {
				grid.Set(p, '0')
			} else {
				anyPNotFlash = true
			}
		})
		if !anyPNotFlash {
			fmt.Println(step)
			break
		}
	}
}

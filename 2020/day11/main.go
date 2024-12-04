package main

import (
	"AOC/h"
	"fmt"
	"slices"
)

func main() {
	fmt.Println("Start")
	// part1()
	part2()
}

const (
	FLOOR          byte = '.'
	EMTYSEAT       byte = 'L'
	FULLSEAT       byte = '#'
	MARKEDForFull  byte = 'M'
	MARKEDForEmpty byte = 'E'
)

func part1() {
	sum := 0
	lines := h.GetLinesAsSlice()
	grid := h.ConvertLinesToGrid(lines)
	hash := grid.HashValue()

	for {
		grid.ForEachPoint(func(p h.Point) {
			if grid.At(p) == EMTYSEAT && !slices.Contains(grid.Neighbours(p), FULLSEAT) {
				grid.Set(p, MARKEDForFull)
			}
			if grid.At(p) == FULLSEAT && len(h.Filter(grid.Neighbours(p), func(e byte) bool { return e == FULLSEAT || e == MARKEDForEmpty })) >= 4 {
				grid.Set(p, MARKEDForEmpty)
			}
		})

		grid.ForEachPoint(func(p h.Point) {
			if grid.At(p) == MARKEDForFull {
				grid.Set(p, FULLSEAT)
			}
			if grid.At(p) == MARKEDForEmpty {
				grid.Set(p, EMTYSEAT)
			}
		})

		newHash := grid.HashValue()

		//fmt.Println("----------")
		//h.PrintGrid(grid)
		//fmt.Println("sum", sum)
		//fmt.Println("hash", newHash)
		//fmt.Println("----------")
		//fmt.Println("")

		if sum%100 == 0 {
			fmt.Println("Runde", sum)
		}

		if newHash == hash {
			break
		}
		hash = newHash
		sum++
	}
	count := 0
	grid.ForEachPoint(func(p h.Point) {
		if grid.At(p) == FULLSEAT {
			count++
		}
	})
	fmt.Println("count", count)

}

func part2() {
	sum := 0
	lines := h.GetLinesAsSlice()
	grid := h.ConvertLinesToGrid(lines)
	hash := grid.HashValue()

	for {
		grid.ForEachPoint(func(p h.Point) {

			if grid.At(p) == EMTYSEAT {
				seeFull := false
				for _, dir := range h.GetAllDirs() {
					h.WalkThroughLine(grid, dir, p, func(pos h.Point, i int) bool {
						if p == pos {
							return true
						}
						if grid.At(pos) == FULLSEAT {
							seeFull = true
							return false
						}
						if grid.At(pos) == EMTYSEAT || grid.At(pos) == MARKEDForFull {
							return false
						}
						return true
					})
				}
				if !seeFull {
					grid.Set(p, MARKEDForFull)
				}
			}

			if grid.At(p) == FULLSEAT {
				numOfFullOrMarkedFE := 0
				for _, dir := range h.GetAllDirs() {
					h.WalkThroughLine(grid, dir, p, func(pos h.Point, i int) bool {
						if pos == p {
							return true
						}
						if grid.At(pos) == FULLSEAT || grid.At(pos) == MARKEDForEmpty {
							numOfFullOrMarkedFE++
							return false
						}
						if grid.At(pos) == EMTYSEAT {
							return false
						}
						return true
					})
				}

				if numOfFullOrMarkedFE >= 5 {
					grid.Set(p, MARKEDForEmpty)
				}
			}

		})

		grid.ForEachPoint(func(p h.Point) {
			if grid.At(p) == MARKEDForFull {
				grid.Set(p, FULLSEAT)
			}
			if grid.At(p) == MARKEDForEmpty {
				grid.Set(p, EMTYSEAT)
			}
		})

		newHash := grid.HashValue()

		//fmt.Println("----------")
		//h.PrintGrid(grid)
		//fmt.Println("sum", sum)
		//fmt.Println("hash", newHash)
		//fmt.Println("----------")
		//fmt.Println("")

		if sum%100 == 0 {
			fmt.Println("Runde", sum)
		}

		if newHash == hash {
			break
		}
		hash = newHash
		sum++
	}
	count := 0
	grid.ForEachPoint(func(p h.Point) {
		if grid.At(p) == FULLSEAT {
			count++
		}
	})
	fmt.Println("count", count)
}

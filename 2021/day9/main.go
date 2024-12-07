package main

import (
	"AOC/h"
	"fmt"
	"slices"
)

func main() {
	fmt.Println("Start")
	//part1()
	part2()
}

func part1() {
	sum := 0
	lines := h.GetLinesAsSlice()
	grid := h.ConvertLinesToGrid(lines)
	lows := make([]int, 0)
	grid.ForEachPoint(func(p h.Point) {
		isLow := true
		cur := grid.AtNum(p)
		for _, dir := range h.GetBasicDirs() {
			pos := p.Relative(dir)
			if pos.IsInGrid(grid) {
				r := grid.AtNum(pos)
				if cur >= r {
					isLow = false
				}
			}
		}
		if isLow {
			lows = append(lows, cur)
		}
	})

	for _, p := range lows {
		sum += p + 1
	}
	fmt.Println(lows)
	fmt.Println(sum)
}

func part2() {
	sum := 1
	lines := h.GetLinesAsSlice()
	grid := h.ConvertLinesToGrid(lines)
	lows := make([]h.Point, 0)
	grid.ForEachPoint(func(p h.Point) {
		isLow := true
		cur := grid.AtNum(p)
		for _, dir := range h.GetBasicDirs() {
			pos := p.Relative(dir)
			if pos.IsInGrid(grid) {
				r := grid.AtNum(pos)
				if cur >= r {
					isLow = false
				}
			}
		}
		if isLow {
			lows = append(lows, p)
		}
	})

	amountChanged := make([]int, 0)
	for _, l := range lows {
		amountChanged = append(amountChanged, grid.FloodFillBasic(l, byte('X'), func(p h.Point) bool {
			if grid.At(p) == '9' {
				return true
			}
			return false
		}))
	}
	slices.Sort(amountChanged)

	for i := len(amountChanged) - 1; i > len(amountChanged)-4; i-- {
		sum *= amountChanged[i]
	}
	fmt.Println(sum)
}

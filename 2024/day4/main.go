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

	h.WalkThrough(grid, h.EAST, h.Point{}, func(p h.Point) {
		res := make([]bool, 8)
		for i, _ := range res {
			res[i] = false
		}

		res[0] = grid.GridCompareByteArr(p, h.EAST, []byte("XMAS"))
		res[1] = grid.GridCompareByteArr(p, h.NORTH, []byte("XMAS"))
		res[2] = grid.GridCompareByteArr(p, h.NORTHEAST, []byte("XMAS"))
		res[3] = grid.GridCompareByteArr(p, h.SOUTHEAST, []byte("XMAS"))
		res[4] = grid.GridCompareByteArr(p, h.WEST, []byte("XMAS"))
		res[5] = grid.GridCompareByteArr(p, h.SOUTH, []byte("XMAS"))
		res[6] = grid.GridCompareByteArr(p, h.NORTHWEST, []byte("XMAS"))
		res[7] = grid.GridCompareByteArr(p, h.SOUTHWEST, []byte("XMAS"))

		for _, b := range res {
			if b {
				sum++
			}
		}
	})

	fmt.Println(sum)
}

func part2() {
	sum := 0
	lines := h.GetLinesAsSlice()
	grid := h.ConvertLinesToGrid(lines)

	h.WalkThrough(grid, h.EAST, h.Point{}, func(p h.Point) {
		res := make([]bool, 4)
		for i, _ := range res {
			res[i] = false
		}

		res[0] = h.GridCompareByteArr(grid, p, h.NORTHEAST, []byte("AS")) && h.GridCompareByteArr(grid, p, h.SOUTHWEST, []byte("AM"))
		res[1] = h.GridCompareByteArr(grid, p, h.NORTHEAST, []byte("AM")) && h.GridCompareByteArr(grid, p, h.SOUTHWEST, []byte("AS"))
		res[2] = h.GridCompareByteArr(grid, p, h.SOUTHEAST, []byte("AS")) && h.GridCompareByteArr(grid, p, h.NORTHWEST, []byte("AM"))
		res[3] = h.GridCompareByteArr(grid, p, h.SOUTHEAST, []byte("AM")) && h.GridCompareByteArr(grid, p, h.NORTHWEST, []byte("AS"))

		fmt.Println(res)

		if (res[0] || res[1]) && (res[2] || res[3]) {
			sum++
		}
	})

	fmt.Println(sum)
}

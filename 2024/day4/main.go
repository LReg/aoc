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

	h.WalkThrough(grid, h.HORIZONTAL, h.Point{}, func(p h.Point) {
		res := make([]bool, 8)
		for i, _ := range res {
			res[i] = false
		}

		res[0] = h.GridCompareByte(grid, p, h.HORIZONTAL, []byte("XMAS"))
		res[1] = h.GridCompareByte(grid, p, h.VERTICAL, []byte("XMAS"))
		res[2] = h.GridCompareByte(grid, p, h.DIAGONALUP, []byte("XMAS"))
		res[3] = h.GridCompareByte(grid, p, h.DIAGONALDOWN, []byte("XMAS"))
		res[4] = h.GridCompareByte(grid, p, h.HORIZONTALREVERSE, []byte("XMAS"))
		res[5] = h.GridCompareByte(grid, p, h.VERTICALREVERSE, []byte("XMAS"))
		res[6] = h.GridCompareByte(grid, p, h.DIAGONALUPREVERSE, []byte("XMAS"))
		res[7] = h.GridCompareByte(grid, p, h.DIAGONALDOWNREVERSE, []byte("XMAS"))

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

	h.WalkThrough(grid, h.HORIZONTAL, h.Point{}, func(p h.Point) {
		res := make([]bool, 4)
		for i, _ := range res {
			res[i] = false
		}

		res[0] = h.GridCompareByte(grid, p, h.DIAGONALUP, []byte("AS")) && h.GridCompareByte(grid, p, h.DIAGONALDOWNREVERSE, []byte("AM"))
		res[1] = h.GridCompareByte(grid, p, h.DIAGONALUP, []byte("AM")) && h.GridCompareByte(grid, p, h.DIAGONALDOWNREVERSE, []byte("AS"))
		res[2] = h.GridCompareByte(grid, p, h.DIAGONALDOWN, []byte("AS")) && h.GridCompareByte(grid, p, h.DIAGONALUPREVERSE, []byte("AM"))
		res[3] = h.GridCompareByte(grid, p, h.DIAGONALDOWN, []byte("AM")) && h.GridCompareByte(grid, p, h.DIAGONALUPREVERSE, []byte("AS"))

		fmt.Println(res)

		if (res[0] || res[1]) && (res[2] || res[3]) {
			sum++
		}
	})

	fmt.Println(sum)
}

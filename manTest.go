package main

import (
	"AOC/h"
	"fmt"
)

func main() {
	sum := 0
	lines := h.GetLinesAsSlice()
	grid := h.ConvertLinesToGrid(lines)
	p := h.Point{3, 5}

	res := h.GridCompareByte(grid, p, h.HORIZONTALREVERSE, []byte("XMAS"))
	fmt.Println("res", res)

	fmt.Println("sum", sum)
}

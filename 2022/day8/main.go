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

	fmt.Println(sum)
}

func part2() {

}
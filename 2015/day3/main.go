package main

import (
	"AOC/h"
	"fmt"
	"slices"
)

func main() {
	fmt.Println("Start")
	part1()
	part2()
}

func tc(c byte) int {
	switch c {
	case '^':
		return h.NORTH
	case '>':
		return h.EAST
	case 'v':
		return h.SOUTH
	case '<':
		return h.WEST
	}
	return -1
}

func part1() {
	sum := 1
	lines := h.GetLinesAsSlice()
	position := h.Point{0, 0}
	path := []h.Point{position}
	for _, c := range lines[0] {
		dir := tc(byte(c))
		position = position.Relative(dir)
		if !slices.Contains(path, position) {
			sum++
		}
		path = append(path, position)
	}

	fmt.Println(sum)
}

func part2() {
	sum := 1
	lines := h.GetLinesAsSlice()
	position := h.Point{0, 0}
	positionRS := h.Point{0, 0}
	path := []h.Point{position}
	for i, c := range lines[0] {
		isRS := i%2 == 1
		dir := tc(byte(c))
		if !isRS {
			position = position.Relative(dir)
			if !slices.Contains(path, position) {
				sum++
			}
			path = append(path, position)
		} else {
			positionRS = positionRS.Relative(dir)
			if !slices.Contains(path, positionRS) {
				sum++
			}
			path = append(path, positionRS)
		}
	}

	fmt.Println(sum)
}

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

func translateDir(d string) int {
	switch d {
	case "N":
		return h.NORTH
	case "S":
		return h.SOUTH
	case "E":
		return h.EAST
	case "W":
		return h.WEST
	case "L":
		return -1
	case "R":
		return -2
	case "F":
		return -3
	}
	return -4
}

// left == false ; right == true
func turnByDeg(dir int, deg int, turnDir bool) int {
	times := deg / 90
	for i := 0; i < times; i++ {
		if turnDir {
			switch dir {
			case h.NORTH:
				dir = h.EAST
			case h.EAST:
				dir = h.SOUTH
			case h.SOUTH:
				dir = h.WEST
			case h.WEST:
				dir = h.NORTH
			default:
			}
		} else {
			switch dir {
			case h.NORTH:
				dir = h.WEST
			case h.WEST:
				dir = h.SOUTH
			case h.SOUTH:
				dir = h.EAST
			case h.EAST:
				dir = h.NORTH
			default:
			}
		}
	}
	return dir
}

func part1() {
	lines := h.GetLinesAsSlice()
	startPos := h.Point{0, 0}
	currentPos := h.Point{0, 0}
	direction := h.EAST
	for _, l := range lines {
		amount := h.ExtrapolateNumbersFromStringIgnoreNonDig(l)[0]
		d := translateDir(string(l[0]))
		if d >= 0 {
			currentPos = currentPos.RelativeN(d, amount)
		} else {
			if d == -2 {
				direction = turnByDeg(direction, amount, true)
			} else if d == -1 {
				direction = turnByDeg(direction, amount, false)
			} else if d == -3 {
				currentPos = currentPos.RelativeN(direction, amount)
			}
		}
	}
	fmt.Println(startPos.ManhattanDist(currentPos))
}

func part2() {
	lines := h.GetLinesAsSlice()
	startPos := h.Point{0, 0}
	currentPos := h.Point{0, 0}
	waypoint := h.Point{10, 1}
	direction := h.EAST
	for _, l := range lines {
		amount := h.ExtrapolateNumbersFromStringIgnoreNonDig(l)[0]
		d := translateDir(string(l[0]))
		if d >= 0 {
			currentPos = currentPos.RelativeN(d, amount)
		} else {
			if d == -2 {
				direction = turnByDeg(direction, amount, true)
			} else if d == -1 {
				direction = turnByDeg(direction, amount, false)
			} else if d == -3 {
				currentPos = currentPos.RelativeN(direction, amount)
			}
		}
	}
	fmt.Println(startPos.ManhattanDist(currentPos))
}

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

var DIAL_SIZE = 99

func moveRadial(from int, by int) int {
	to := from + by
	if to > DIAL_SIZE || to < 0 {
		to = to % (DIAL_SIZE + 1)
	}
	return to
}

func moveRadialRadialed(from int, by int) (int, int) {
	to := from + by
	radialed := 0
	fmt.Println("to", to)
	if to > DIAL_SIZE {
		radialed = to / (DIAL_SIZE + 1)
		to = to % (DIAL_SIZE + 1)
		return to, radialed
	} else if to < 0 {
		radialed = -to / (DIAL_SIZE + 1)
		radialed++
		to = to % (DIAL_SIZE + 1)
		to = to + DIAL_SIZE + 1
		to = to % (DIAL_SIZE + 1)
		return to, radialed
	}
	if to == 0 {
		radialed++
	}
	return to, radialed
}

func part1() {
	dial := 50
	sum := 0
	lines := h.GetLinesAsSlice()

	for _, line := range lines {
		var LEFT bool
		if line[0] == 'L' {
			LEFT = true
		} else {
			LEFT = false
		}

		moveBy := h.ExtrapolateNumbersFromStringIgnoreNonDig(line)[0]
		if LEFT {
			moveBy = -moveBy
		}

		dial = moveRadial(dial, moveBy)

		if dial == 0 {
			sum++
		}
	}

	fmt.Println(sum)
}

func part2() {
	dial := 50
	sum := 0
	lines := h.GetLinesAsSlice()

	for _, line := range lines {
		var LEFT bool
		if line[0] == 'L' {
			LEFT = true
		} else {
			LEFT = false
		}

		moveBy := h.ExtrapolateNumbersFromStringIgnoreNonDig(line)[0]
		if LEFT {
			moveBy = -moveBy
		}

		dialNew, radialed := moveRadialRadialed(dial, moveBy)
		fmt.Println("newDial, radialed", dialNew, radialed, "\n")
		dial = dialNew

		if radialed > 0 {
			sum += radialed
		}
	}

	fmt.Println(sum)
}

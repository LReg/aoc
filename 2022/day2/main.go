package main

import (
	"AOC/h"
	"fmt"
	"strings"
)

const (
	WIN = iota
	DRAW
	LOOSE
	ERROR
)

type STATE int

func main() {
	part1()
	part2()
}

func part2() {
	lines := h.GetLinesAsSlice()
	sum := 0
	for _, l := range lines {
		game := strings.Split(l, " ")

		var ws int
		var state STATE
		if game[1] == "X" {
			ws = winScore(LOOSE)
			state = LOOSE
		} else if game[1] == "Y" {
			ws = winScore(DRAW)
			state = DRAW
		} else {
			ws = winScore(WIN)
			state = WIN
		}

		s := depCardScore(state, game[0])

		gamePoints := s + ws
		sum += gamePoints
	}
	fmt.Println(sum)
}

func depCardScore(state STATE, enemy string) int {
	if state == DRAW {
		return score(enemy)
	}
	if state == WIN {
		if enemy == "A" {
			return score("B")
		} else if enemy == "B" {
			return score("C")
		} else {
			return score("A")
		}
	}
	if state == LOOSE {
		if enemy == "A" {
			return score("C")
		} else if enemy == "B" {
			return score("A")
		} else {
			return score("B")
		}
	}
	return 0
}

func part1() {
	lines := h.GetLinesAsSlice()
	sum := 0
	for _, l := range lines {
		game := strings.Split(l, " ")
		r := rune(game[1][0])
		r = r - ('X' - 'A')
		game[1] = string(r)

		s := score(game[1])
		win := isWin(game)
		ws := winScore(win)
		if win == ERROR {
			panic("invalid isWin")
		}
		gamePoints := s + ws
		sum += gamePoints
	}
	fmt.Println(sum)
}

func winScore(win STATE) int {
	switch win {
	case WIN:
		return 6
	case LOOSE:
		return 0
	case DRAW:
		return 3
	}
	return -1
}

func isWin(game []string) STATE {

	if game[0] == game[1] {
		return DRAW
	}

	if game[0] == "A" {
		if game[1] == "B" {
			return WIN
		} else {
			return LOOSE
		}
	}
	if game[0] == "B" {
		if game[1] == "A" {
			return LOOSE
		} else {
			return WIN
		}
	}
	if game[0] == "C" {
		if game[1] == "A" {
			return WIN
		} else {
			return LOOSE
		}
	}
	return ERROR
}

func score(s string) int {
	switch s {
	case "A":
		return 1
	case "B":
		return 2
	case "C":
		return 3
	}
	return 0
}

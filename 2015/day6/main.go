package main

import (
	"AOC/h"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Start")
	part1()
	part2()
}

func part1() {
	sum := 0
	lines := h.GetLinesAsSlice()
	grid := h.CreateGrid(1000, 1000)
	grid.ForEachPoint(func(p h.Point) {
		grid.Set(p, 'o')
	})
	for _, l := range lines {
		p := h.ExtrapolateNumbersFromStringIgnoreNonDig(l)
		turnOff := strings.Contains(l, "turn off")
		turnOn := strings.Contains(l, "turn on")
		toggle := strings.Contains(l, "toggle")
		for x := p[0]; x <= p[2]; x++ {
			for y := p[1]; y <= p[3]; y++ {
				if turnOff {
					grid.Set(h.Point{x, y}, 'o')
				} else if turnOn {
					grid.Set(h.Point{x, y}, 'O')
				} else if toggle {
					c := grid.At(h.Point{x, y})
					if c == 'o' {
						grid.Set(h.Point{x, y}, 'O')
					} else {
						grid.Set(h.Point{x, y}, 'o')
					}
				}
			}
		}
	}
	grid.ForEachPoint(func(p h.Point) {
		if grid.At(p) == 'O' {
			sum++
		}
	})
	fmt.Println(sum)
}

func part2() {
	sum := 0
	lines := h.GetLinesAsSlice()
	grid := h.CreateOrderedGrid[int](1000, 1000)
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			grid[x][y] = 0
		}
	}
	for _, l := range lines {
		p := h.ExtrapolateNumbersFromStringIgnoreNonDig(l)
		turnOff := strings.Contains(l, "turn off")
		turnOn := strings.Contains(l, "turn on")
		toggle := strings.Contains(l, "toggle")
		for x := p[0]; x <= p[2]; x++ {
			for y := p[1]; y <= p[3]; y++ {
				if turnOff {
					c := grid[x][y]
					grid[x][y] = h.Max(c-1, 0)
				} else if turnOn {
					c := grid[x][y]
					grid[x][y] = c + 1
				} else if toggle {
					c := grid[x][y]
					grid[x][y] = c + 2
				}
			}
		}
	}
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			sum += grid[x][y]
		}
	}
	fmt.Println(sum)
}

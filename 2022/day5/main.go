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
	lines := h.GetLinesAsSlice()
	endIndex := 0
	for i, l := range lines {
		if strings.Index(l, "1") != -1 {
			endIndex = i
			break
		}
	}
	grid := h.ConvertLinesToGrid(lines[:endIndex])

	// grid to stacks
	// stack by staack
	stacks := make([][]byte, 0)
	for x := 1; x < len(grid); x += 4 {
		stack := make([]byte, 0)
		for y := len(grid[0]) - 1; y >= 0; y-- {
			var b byte
			if grid[x][y] != b && grid[x][y] != byte(' ') {
				stack = append(stack, grid[x][y])
			}
		}
		stacks = append(stacks, stack)
	}

	commandStart := endIndex + 2
	for i := commandStart; i < len(lines); i++ {
		com := h.ExtrapolateNumbersFromStringIgnoreNonDig(lines[i])
		from := com[1] - 1
		to := com[2] - 1
		for _, _ = range h.Iter(com[0]) {
			c := stacks[from][len(stacks[from])-1]
			stacks[from] = stacks[from][:len(stacks[from])-1]
			stacks[to] = append(stacks[to], c)
		}
	}

	for _, s := range stacks {
		fmt.Printf("%c", s[len(s)-1])
	}
	println()

}

func printStacks(stacks [][]byte) {
	fmt.Println("----Stack----")
	for _, s := range stacks {
		for _, c := range s {
			fmt.Printf("%c", c)
		}
		fmt.Printf("\n")
	}
	fmt.Println("-----------")
}

func part2() {
	lines := h.GetLinesAsSlice()
	endIndex := 0
	for i, l := range lines {
		if strings.Index(l, "1") != -1 {
			endIndex = i
			break
		}
	}
	grid := h.ConvertLinesToGrid(lines[:endIndex])

	// grid to stacks
	// stack by staack
	stacks := make([][]byte, 0)
	for x := 1; x < len(grid); x += 4 {
		stack := make([]byte, 0)
		for y := len(grid[0]) - 1; y >= 0; y-- {
			var b byte
			if grid[x][y] != b && grid[x][y] != byte(' ') {
				stack = append(stack, grid[x][y])
			}
		}
		stacks = append(stacks, stack)
	}

	commandStart := endIndex + 2
	for i := commandStart; i < len(lines); i++ {
		com := h.ExtrapolateNumbersFromStringIgnoreNonDig(lines[i])
		amount := com[0]
		from := com[1] - 1
		to := com[2] - 1
		c := stacks[from][len(stacks[from])-amount : len(stacks[from])]
		stacks[to] = append(stacks[to], c...)
		stacks[from] = stacks[from][:len(stacks[from])-amount]
	}

	for _, s := range stacks {
		fmt.Printf("%c", s[len(s)-1])
	}
	println()
}

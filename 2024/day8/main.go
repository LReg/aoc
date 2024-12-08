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

func part1() {
	lines := h.GetLinesAsSlice()
	grid := h.ConvertLinesToGrid(lines)

	frequencies := make([]byte, 0)
	grid.ForEachPoint(func(p h.Point) {
		b := grid.At(p)
		if !slices.Contains(frequencies, b) && b != '.' {
			frequencies = append(frequencies, b)
		}
	})

	antinodes := make([]h.Point, 0)
	for _, f := range frequencies {
		positions := make([]h.Point, 0)
		grid.ForEachPoint(func(p h.Point) {
			if grid.At(p) == f {
				positions = append(positions, p)
			}
		})

		combinations := h.CrossProduct(positions, 2)
		combinations = h.Filter(combinations, func(c []h.Point) bool {
			if c[0] == c[1] {
				return false
			} else {
				return true
			}
		})

		for _, comb := range combinations {
			xDiff := comb[0].X - comb[1].X
			yDiff := comb[0].Y - comb[1].Y
			antinode := h.Point{
				comb[0].X + xDiff,
				comb[0].Y + yDiff,
			}
			if antinode.IsInGrid(grid) && !slices.Contains(antinodes, antinode) {
				antinodes = append(antinodes, antinode)
			}
		}
	}

	fmt.Println(len(antinodes))
}

func part2() {
	lines := h.GetLinesAsSlice()
	grid := h.ConvertLinesToGrid(lines)

	frequencies := make([]byte, 0)
	grid.ForEachPoint(func(p h.Point) {
		b := grid.At(p)
		if !slices.Contains(frequencies, b) && b != '.' {
			frequencies = append(frequencies, b)
		}
	})

	antinodes := make([]h.Point, 0)
	for _, f := range frequencies {
		positions := make([]h.Point, 0)
		grid.ForEachPoint(func(p h.Point) {
			if grid.At(p) == f {
				positions = append(positions, p)
			}
		})

		combinations := h.CrossProduct(positions, 2)
		combinations = h.Filter(combinations, func(c []h.Point) bool {
			if c[0] == c[1] {
				return false
			} else {
				return true
			}
		})

		for _, comb := range combinations {
			xDiff := comb[0].X - comb[1].X
			yDiff := comb[0].Y - comb[1].Y
			for {
				antinode := h.Point{
					comb[0].X + xDiff,
					comb[0].Y + yDiff,
				}
				comb[0].X += xDiff
				comb[0].Y += yDiff
				if antinode.IsInGrid(grid) && !slices.Contains(antinodes, antinode) && grid.At(antinode) == '.' {
					antinodes = append(antinodes, antinode)
				}
				if !antinode.IsInGrid(grid) {
					break
				}
			}

		}
	}

	antinodesGrid := grid.Copy()
	antinodesGrid.ForEachPoint(func(p h.Point) {
		if antinodesGrid.At(p) == '.' && slices.Contains(antinodes, p) {
			antinodesGrid.Set(p, '#')
		}
	})
	nrOfAntennas := 0
	grid.ForEachPoint(func(p h.Point) {
		if grid.At(p) != '.' {
			nrOfAntennas++
		}
	})
	h.PrintGrid(antinodesGrid)
	fmt.Println("antinodes len", len(antinodes))
	fmt.Println("antennas nr", nrOfAntennas)

	fmt.Println(len(antinodes) + nrOfAntennas)
}

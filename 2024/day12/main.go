package main

import (
	"AOC/h"
	"fmt"
	"slices"
	"strconv"
)

func main() {
	fmt.Println("Start")
	//part1()
	//part2()
	part2thathopefullyworks()
}

func hasPlotLeft(grid h.Grid) (bool, byte, h.Point) {
	hasLeft := false
	ch := byte(' ')
	po := h.Point{}
	grid.ForEachPoint(func(p h.Point) {
		c := grid.At(p)
		if c != '#' && !hasLeft {
			hasLeft = true
			ch = c
			po = p
		}
	})
	return hasLeft, ch, po
}

type Plot struct {
	area      int
	perimeter int
}

func part1() {
	sum := 0
	lines := h.GetLinesAsSlice()
	grid := h.ConvertLinesToGrid(lines)
	plots := []Plot{}
	for {
		has, ch, point := hasPlotLeft(grid)
		if !has {
			break
		}
		area := grid.FloodFillBasic(point, '?', func(p h.Point) bool { return grid.At(p) != ch })

		perimeter := 0
		copyGrid := h.CreateGrid(len(grid)+2, len(grid[0])+2)
		copyGrid.ForEachPoint(func(p h.Point) {
			copyGrid.Set(p, '+')
		})

		grid.ForEachPoint(func(p h.Point) {
			if p.IsInGrid(grid) {
				copyGrid.Set(h.Point{p.X + 1, p.Y + 1}, grid.At(p))
			}
		})

		copyGrid.ForEachPoint(func(p h.Point) {
			for _, n := range p.BasicNeighbours() {
				if n.IsInGrid(copyGrid) && copyGrid.At(n) == '?' && copyGrid.At(p) != '?' {
					perimeter++
				}
			}
		})

		fmt.Println(Plot{area, perimeter})

		plots = append(plots, Plot{area, perimeter})

		grid.ForEachPoint(func(p h.Point) {
			if grid.At(p) == '?' {
				grid.Set(p, '#')
			}
		})
	}
	for _, p := range plots {
		sum += p.area * p.perimeter
	}
	fmt.Println(sum)
}

// does not work
func part2() {
	sum := 0
	lines := h.GetLinesAsSlice()
	grid := h.ConvertLinesToGrid(lines)
	plots := []Plot{}
	for {
		has, ch, point := hasPlotLeft(grid)
		if !has {
			break
		}
		area := grid.FloodFillBasic(point, '?', func(p h.Point) bool { return grid.At(p) != ch })

		copyGrid := h.CreateGrid(len(grid)+2, len(grid[0])+2)
		copyGrid.ForEachPoint(func(p h.Point) {
			copyGrid.Set(p, '+')
		})

		grid.ForEachPoint(func(p h.Point) {
			if p.IsInGrid(grid) {
				copyGrid.Set(h.Point{p.X + 1, p.Y + 1}, grid.At(p))
			}
		})

		startPointMinus := h.Point{}
		copyGrid.ForEachPoint(func(p h.Point) {
			for _, n := range p.BasicNeighbours() {
				if n.IsInGrid(copyGrid) && copyGrid.At(n) == '?' && copyGrid.At(p) != '?' {
					copyGrid.Set(p, '-')
				}
			}
		})
		copyGrid.ForEachPoint(func(p h.Point) {
			count := 0
			for _, n := range p.BasicNeighbours() {
				if n.IsInGrid(copyGrid) && copyGrid.At(n) == '-' {
					count++
				}
			}
			countQ := 0
			for _, n := range p.Neighbours() {
				if n.IsInGrid(copyGrid) && copyGrid.At(n) == '?' {
					countQ++
				}
			}
			if count >= 2 && copyGrid.At(p) != '?' && countQ >= 1 {
				copyGrid.Set(p, '-')
				startPointMinus = p
			}
		})

		sides := 0
		h.PrintGrid(copyGrid)
		for {
			horizontal := 0
			stillContainsMinus := false
			for {
				hadN := false
				for _, n := range startPointMinus.BasicNeighbours() {
					if n.IsInGrid(copyGrid) && copyGrid.At(n) == '-' {
						if (startPointMinus.Relative(h.WEST) == n || startPointMinus.Relative(h.EAST) == n) && (horizontal == 0 || horizontal == 2) {
							sides++
							horizontal = 1
						} else if (startPointMinus.Relative(h.NORTH) == n || startPointMinus.Relative(h.SOUTH) == n) && (horizontal == 0 || horizontal == 1) {
							sides++
							horizontal = 2
						}
						copyGrid.Set(startPointMinus, '#')
						startPointMinus = n
						hadN = true
						break
					}
				}
				if !hadN {
					copyGrid.Set(startPointMinus, '#')
					break
				}

				h.PrintGrid(copyGrid)
				fmt.Println(sides)
			}
			copyGrid.ForEachPoint(func(p h.Point) {
				if copyGrid.At(p) == '-' {
					stillContainsMinus = true
					startPointMinus = p
				}
			})

			if !stillContainsMinus {
				break
			}
		}

		plots = append(plots, Plot{area, sides})

		grid.ForEachPoint(func(p h.Point) {
			if grid.At(p) == '?' {
				grid.Set(p, '#')
			}
		})
	}
	for _, p := range plots {
		sum += p.area * p.perimeter
	}
	fmt.Println(sum)
}

// all testinputs work but still not working
func part2thathopefullyworks() {
	sum := 0
	lines := h.GetLinesAsSlice()
	grid := h.ConvertLinesToGrid(lines)
	plots := []Plot{}
	for {
		has, ch, point := hasPlotLeft(grid)
		if !has {
			break
		}
		area := grid.FloodFillBasic(point, '?', func(p h.Point) bool { return grid.At(p) != ch })

		copyGrid := h.CreateGrid(len(grid)+2, len(grid[0])+2)
		copycopyGrid := h.CreateGrid(len(grid)+2, len(grid[0])+2)
		copyGrid.ForEachPoint(func(p h.Point) {
			copyGrid.Set(p, '+')
			copycopyGrid.Set(p, '+')

		})

		grid.ForEachPoint(func(p h.Point) {
			if p.IsInGrid(grid) {
				copyGrid.Set(h.Point{p.X + 1, p.Y + 1}, grid.At(p))
			}
		})

		copyGrid.ForEachPoint(func(p h.Point) {
			for _, n := range p.BasicNeighbours() {
				if n.IsInGrid(copyGrid) && copyGrid.At(n) == '?' && copyGrid.At(p) != '?' {
					copyGrid.Set(p, '-')
				}
			}
		})
		copyGrid.ForEachPoint(func(p h.Point) {
			count := 0
			for _, n := range p.BasicNeighbours() {
				if n.IsInGrid(copyGrid) && copyGrid.At(n) == '-' {
					count++
				}
			}
			countQ := 0
			for _, n := range p.Neighbours() {
				if n.IsInGrid(copyGrid) && copyGrid.At(n) == '?' {
					countQ++
				}
			}
			if count >= 2 && copyGrid.At(p) != '?' && countQ >= 1 {
				copyGrid.Set(p, '-')
			}
		})

		minuses := []h.Point{}

		copyGrid.ForEachPoint(func(p h.Point) {
			if copyGrid.At(p) == '-' {
				minuses = append(minuses, p)
			}
		})

		counterSum := 0
		for _, m := range minuses {
			counter := 0
			isCornerPiece := false
			nc := 0
			for _, n := range m.BasicNeighbours() {
				if n.IsInGrid(copyGrid) && slices.Contains(minuses, n) {
					nc++
				}
			}
			if ((m.Relative(h.NORTH).IsInGrid(copyGrid)) && copyGrid.At(m.Relative(h.NORTH)) == '-' ||
				(m.Relative(h.SOUTH).IsInGrid(copyGrid)) && copyGrid.At(m.Relative(h.SOUTH)) == '-') &&
				((m.Relative(h.WEST).IsInGrid(copyGrid)) && copyGrid.At(m.Relative(h.WEST)) == '-' ||
					(m.Relative(h.EAST).IsInGrid(copyGrid)) && copyGrid.At(m.Relative(h.EAST)) == '-') &&
				nc == 2 {
				isCornerPiece = true
			}

			isEntrancePeace := false
			isDeadEnd := false

			if nc == 3 {
				isEntrancePeace = true
			}
			if nc == 1 {
				isDeadEnd = true
			}

			if isCornerPiece {
				counter++
			}
			if isEntrancePeace {
				//counter += 1
			}
			if isDeadEnd {
				counter += 4
				for _, dir := range h.GetBasicDirs() {
					pos := m
					for {
						r := pos.Relative(dir)
						if !slices.Contains(minuses, r) {
							break
						}
						neiC := 0
						for _, n := range r.BasicNeighbours() {
							if slices.Contains(minuses, n) {
								neiC++
							}
						}
						if neiC == 3 {
							if slices.Contains(minuses, r.Relative(dir)) {
								counter--
								break
							}
						}
						pos = r
					}
				}
			}
			counterSum += counter
			copycopyGrid.Set(m, strconv.Itoa(counter)[0])
		}

		//fmt.Println(string(ch), area, counterSum, area*counterSum)

		plots = append(plots, Plot{area, counterSum})
		grid.ForEachPoint(func(p h.Point) {
			if grid.At(p) == '?' {
				grid.Set(p, '#')
			}
		})
	}
	for _, p := range plots {
		sum += p.area * p.perimeter
	}
	fmt.Println(sum)
}

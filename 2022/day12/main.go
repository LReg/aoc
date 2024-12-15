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
	lines := h.GetLinesAsSlice()
	grid := h.ConvertLinesToGrid(lines)
	neighbours := h.NewNeighbourMap[h.Point]()
	var start h.Point
	var end h.Point
	grid.ForEachPoint(func(p h.Point) {
		if grid.At(p) == 'S' {
			start = p
		}
		if grid.At(p) == 'E' {
			end = p
		}
	})
	grid.ForEachPoint(func(p h.Point) {
		height := int(grid.At(p))
		if start == p {
			height = int('a')
		} else if end == p {
			height = int('z')
		}

		for _, n := range p.BasicNeighbours() {
			if n.IsInGrid(grid) {
				heightN := int(grid.At(n))

				if start == n {
					heightN = int('a')
				} else if end == n {
					heightN = int('z')
				}

				if height+1 >= heightN {
					neighbours.AddEdge(p, h.NewEdge(n, 1))
				}
			}
		}
	})

	_, l := h.Dijkstra(neighbours, start, end)
	fmt.Println("part 1:", l)
}

func part2() {
	lines := h.GetLinesAsSlice()
	grid := h.ConvertLinesToGrid(lines)
	neighbours := h.NewNeighbourMap[h.Point]()
	startinPoints := []h.Point{}
	var end h.Point
	grid.ForEachPoint(func(p h.Point) {
		if grid.At(p) == 'a' {
			startinPoints = append(startinPoints, p)
		}
		if grid.At(p) == 'E' {
			end = p
		}
	})
	grid.ForEachPoint(func(p h.Point) {
		height := int(grid.At(p))
		if end == p {
			height = int('z')
		}

		for _, n := range p.BasicNeighbours() {
			if n.IsInGrid(grid) {
				heightN := int(grid.At(n))

				if end == n {
					heightN = int('z')
				}

				if height <= heightN+1 {
					neighbours.AddEdge(p, h.NewEdge(n, 1))
				}
			}
		}
	})

	_, l := h.BFS(neighbours, end, func(p h.Point) bool {
		if grid.At(p) == 'a' {
			return true
		} else {
			return false
		}
	})
	fmt.Println("part 2", l)

}

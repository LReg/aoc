package main

import (
	"AOC/h"
	"fmt"
	"math"
)

func main() {
	fmt.Println("Start")
	part1faster()
	//part1()
	part2()
}

func part1faster() {
	sum := 0
	lines := h.GetLinesAsSlice()
	grid := h.ConvertLinesToGrid(lines)
	all0 := []h.Point{}
	all9 := []h.Point{}
	nei := h.NewNeighbourMap[h.Point]()
	grid.ForEachPoint(func(p h.Point) {
		n := grid.AtNum(p)
		if n == 0 {
			all0 = append(all0, p)
		} else if n == 9 {
			all9 = append(all9, p)
		}

		neis := p.BasicNeighbours()
		for _, ni := range neis {
			if ni.IsInGrid(grid) {
				niNum := grid.AtNum(ni)
				if n+1 == niNum {
					nei.AddEdge(p, h.NewEdge(ni, 1))
				}
			}
		}
	})

	for _, z := range all0 {
		score := 0
		for _, n := range all9 {
			if h.BFSAnyFoundPath(nei, z, n, 0, 10) {
				score++
			}
		}
		sum += score
	}

	fmt.Println(sum)
}

func part1() {
	sum := 0
	lines := h.GetLinesAsSlice()
	grid := h.ConvertLinesToGrid(lines)
	all0 := []h.Point{}
	all9 := []h.Point{}
	nei := h.NewNeighbourMap[h.Point]()
	grid.ForEachPoint(func(p h.Point) {
		n := grid.AtNum(p)
		if n == 0 {
			all0 = append(all0, p)
		} else if n == 9 {
			all9 = append(all9, p)
		}

		neis := p.BasicNeighbours()
		for _, ni := range neis {
			if ni.IsInGrid(grid) {
				niNum := grid.AtNum(ni)
				if n+1 == niNum {
					nei.AddEdge(p, h.NewEdge(ni, 1))
				}
			}
		}
	})

	matrix := h.FloydWarshall(nei)

	for _, z := range all0 {
		score := 0
		for _, n := range all9 {
			if matrix[z][n] != math.MaxInt {
				score++
			}
		}
		sum += score
	}

	fmt.Println(sum)
}

func part2() {
	sum := 0
	lines := h.GetLinesAsSlice()
	grid := h.ConvertLinesToGrid(lines)
	all0 := []h.Point{}
	all9 := []h.Point{}
	nei := h.NewNeighbourMap[h.Point]()
	grid.ForEachPoint(func(p h.Point) {
		n := grid.AtNum(p)
		if n == 0 {
			all0 = append(all0, p)
		} else if n == 9 {
			all9 = append(all9, p)
		}

		neis := p.BasicNeighbours()
		for _, ni := range neis {
			if ni.IsInGrid(grid) {
				niNum := grid.AtNum(ni)
				if n+1 == niNum {
					nei = nei.AddEdge(p, h.NewEdge(ni, 1))
				}
			}
		}
	})

	for _, z := range all0 {
		score := 0
		for _, n := range all9 {
			nFound := h.BFSNrOfPaths(nei, z, n, 0, 10)
			score += nFound
		}
		sum += score
	}

	fmt.Println(sum)
}

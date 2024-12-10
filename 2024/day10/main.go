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
	nei := h.NeighbourMap[h.Point]{}
	grid.ForEachPoint(func(p h.Point) {
		n := grid.AtNum(p)
		if n == 0 {
			all0 = append(all0, p)
		} else if n == 9 {
			all9 = append(all9, p)
		}

		neis := p.BasicNeighbours()
		nei[p] = []h.Edge[h.Point]{}
		for _, ni := range neis {
			if ni.IsInGrid(grid) {
				niNum := grid.AtNum(ni)
				if n+1 == niNum {
					nei[p] = append(nei[p], h.Edge[h.Point]{Weight: 1, To: ni})
				}
			}
		}
	})

	var search func(h.Point, h.Point, int) bool
	search = func(st h.Point, end h.Point, dep int) bool {
		if dep == 10 {
			return false
		}
		if st == end {
			return true
		}
		neis := nei[st]
		anyFound := false
		for _, n := range neis {
			if search(n.To, end, dep+1) {
				anyFound = true
			}
		}
		return anyFound
	}

	for _, z := range all0 {
		score := 0
		for _, n := range all9 {
			if search(z, n, 0) {
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
	nei := h.NeighbourMap[h.Point]{}
	grid.ForEachPoint(func(p h.Point) {
		n := grid.AtNum(p)
		if n == 0 {
			all0 = append(all0, p)
		} else if n == 9 {
			all9 = append(all9, p)
		}

		neis := p.BasicNeighbours()
		nei[p] = []h.Edge[h.Point]{}
		for _, ni := range neis {
			if ni.IsInGrid(grid) {
				niNum := grid.AtNum(ni)
				if n+1 == niNum {
					nei[p] = append(nei[p], h.Edge[h.Point]{Weight: 1, To: ni})
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
	nei := h.NeighbourMap[h.Point]{}
	grid.ForEachPoint(func(p h.Point) {
		n := grid.AtNum(p)
		if n == 0 {
			all0 = append(all0, p)
		} else if n == 9 {
			all9 = append(all9, p)
		}

		neis := p.BasicNeighbours()
		nei[p] = []h.Edge[h.Point]{}
		for _, ni := range neis {
			if ni.IsInGrid(grid) {
				niNum := grid.AtNum(ni)
				if n+1 == niNum {
					nei[p] = append(nei[p], h.Edge[h.Point]{Weight: 1, To: ni})
				}
			}
		}
	})

	var search func(h.Point, h.Point, int) int
	search = func(st h.Point, end h.Point, dep int) int {
		if dep == 10 {
			return 0
		}
		if st == end {
			return 1
		}
		neis := nei[st]
		nFound := 0
		for _, n := range neis {
			nr := search(n.To, end, dep+1)
			nFound += nr
		}
		return nFound
	}

	for _, z := range all0 {
		score := 0
		for _, n := range all9 {
			nFound := search(z, n, 0)
			score += nFound
		}
		sum += score
	}

	fmt.Println(sum)
}

package main

import (
	"AOC/h"
	"fmt"
	"slices"
)

func main() {
	fmt.Println("Start")
	//part1()
	part2()
}

type PointTime struct {
	pos  h.Point
	time int
}

func bytePosContains(bpos []PointTime, pos h.Point, time int) bool {
	return slices.ContainsFunc(bpos, func(bp PointTime) bool {
		if bp.pos == pos && bp.time <= time {
			return true
		}
		return false
	})
}

// max x, y
const (
	HW    = 70
	BYTES = 1024
)

func part1() {
	sum := 0
	lines := h.GetLinesAsSlice()
	bpositions := []PointTime{}
	for i, l := range lines {
		xy := h.ExtrapolateNumbersFromString(l, ",")
		bpositions = append(bpositions, PointTime{
			h.Point{xy[0], xy[1]},
			i,
		})
	}
	path, l := bfs(bpositions)
	grid := h.CreateGrid(HW+1, HW+1)
	grid.ForEachPoint(func(p h.Point) {
		if slices.Contains(path, p) {
			grid.Set(p, 'O')
		} else {
			grid.Set(p, '.')
		}
	})
	h.PrintGrid(grid)
	fmt.Println(path, l)
	fmt.Println(sum)
}

func inBounds(p h.Point) bool {
	return p.X >= 0 && p.Y >= 0 && p.X <= HW && p.Y <= HW
}

func bfs(bpositions []PointTime) ([]h.Point, int) {
	open := []PointTime{{pos: h.Point{0, 0}, time: 0}}
	closed := []h.Point{}
	distance := map[h.Point]int{h.Point{0, 0}: 0}
	prev := map[h.Point]h.Point{}
	end := h.Point{HW, HW}
	for len(open) > 0 {
		next := open[0]
		open = open[1:]
		if slices.Contains(closed, next.pos) {
			continue
		}
		if next.pos == end {
			break
		}
		closed = append(closed, next.pos)
		for _, n := range next.pos.BasicNeighbours() {
			if !inBounds(n) {
				continue
			}
			if bytePosContains(bpositions, n, BYTES-1) {
				continue
			}
			if slices.Contains(closed, n) {
				continue
			}
			distance[n] = distance[next.pos] + 1
			open = append(open, PointTime{n, next.time + 1})
			prev[n] = next.pos
		}
	}
	path := []h.Point{}
	current := end
	st := h.Point{0, 0}
	for current != st {
		path = append(path, current)
		current = prev[current]
	}
	path = append(path, current)
	slices.Reverse(path)
	return path, distance[end]
}

func bfsP2(bpositions []PointTime, bytes int) ([]h.Point, int) {
	open := []PointTime{{pos: h.Point{0, 0}, time: 0}}
	closed := []h.Point{}
	distance := map[h.Point]int{h.Point{0, 0}: 0}
	prev := map[h.Point]h.Point{}
	end := h.Point{HW, HW}
	for len(open) > 0 {
		next := open[0]
		open = open[1:]
		if slices.Contains(closed, next.pos) {
			continue
		}
		if next.pos == end {
			break
		}
		closed = append(closed, next.pos)
		for _, n := range next.pos.BasicNeighbours() {
			if !inBounds(n) {
				continue
			}
			if bytePosContains(bpositions, n, bytes-1) {
				continue
			}
			if slices.Contains(closed, n) {
				continue
			}
			distance[n] = distance[next.pos] + 1
			open = append(open, PointTime{n, next.time + 1})
			prev[n] = next.pos
		}
	}
	path := []h.Point{}
	current := end
	st := h.Point{0, 0}
	for current != st {
		path = append(path, current)
		current = prev[current]
	}
	path = append(path, current)
	slices.Reverse(path)
	return path, distance[end]
}

func part2() {
	sum := 0
	lines := h.GetLinesAsSlice()
	bpositions := []PointTime{}
	for i, l := range lines {
		xy := h.ExtrapolateNumbersFromString(l, ",")
		bpositions = append(bpositions, PointTime{
			h.Point{xy[0], xy[1]},
			i,
		})
	}
	for i := BYTES; true; i++ {
		_, l := bfsP2(bpositions, i)
		if l == 0 {
			fmt.Println(bpositions[i-1])
			break
		} else {
			fmt.Println("found path for", i, "bytes")
		}
	}

	fmt.Println(sum)
}

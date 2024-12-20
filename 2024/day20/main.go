package main

import (
	"AOC/h"
	"cmp"
	"fmt"
	"math"
	"slices"
)

func main() {
	fmt.Println("Start")
	//part1()
	part2()
}

func Dijkstra[T cmp.Ordered | h.Point](neighbourMap map[T][]h.Edge[T], start T, end T) ([]T, int, map[T]int) {
	distances := make(map[T]int)
	for point := range neighbourMap {
		distances[point] = math.MaxInt
	}
	distances[start] = 0
	previos := make(map[T]T)

	unvisited := h.NewPC[T]()
	unvisited.Push(start, 0)
	visited := make(map[T]bool)

	for unvisited.Len() > 0 {
		nearestPoint := unvisited.Pop()

		if nearestPoint == end {
			break
		}

		visited[nearestPoint] = true

		// update neighbour distances and add them to the unvisited list
		for _, neighbour := range neighbourMap[nearestPoint] {
			if visited[neighbour.To] {
				continue
			}
			newDistance := distances[nearestPoint] + neighbour.Weight
			if newDistance < distances[neighbour.To] {
				distances[neighbour.To] = newDistance
				previos[neighbour.To] = nearestPoint
				if !unvisited.Contains(neighbour.To) {
					unvisited.Push(neighbour.To, newDistance)
				} else {
					unvisited.UpdatePriority(neighbour.To, newDistance)
				}
			}
		}
	}

	path := make([]T, 0)
	for point := end; point != start; point = previos[point] {
		path = append(path, point)
	}
	path = append(path, start)
	slices.Reverse(path)

	return path, distances[end], distances
}

type Cheat struct {
	start h.Point
	end   h.Point
	diff  int
}

func part1() {
	sum := 0
	lines := h.GetLinesAsSlice()
	grid := h.ConvertLinesToGrid(lines)
	start := h.Point{}
	end := h.Point{}
	nei := h.NewNeighbourMap[h.Point]()
	grid.ForEachPoint(func(p h.Point) {
		if grid.At(p) == 'S' {
			start = p
		} else if grid.At(p) == 'E' {
			end = p
		}
		if grid.At(p) != '#' {
			for _, n := range p.BasicNeighbours() {
				if n.IsInGrid(grid) && grid.At(n) != '#' {
					nei.AddEdge(p, h.NewEdge(n, 1))
				}
			}
		}
	})
	path, l, distances := Dijkstra(nei, start, end)
	for k, v := range distances {
		distances[k] = l - v
	}
	possibleCheats := []Cheat{}
	for _, p := range path {
		for _, c1 := range p.BasicNeighbours() {
			if c1 == p {
				continue
			}
			for _, c2 := range c1.BasicNeighbours() {
				if c2 == c1 || c2 == p {
					continue
				}
				if !c2.IsInGrid(grid) || grid.At(c2) == '#' {
					continue
				}
				cheat := Cheat{c1, c2, distances[p] - distances[c2] - 2}

				if distances[p] > distances[c2]+2 && !slices.Contains(possibleCheats, cheat) {
					possibleCheats = append(possibleCheats, cheat)
				}
			}
		}
	}
	fmt.Println(possibleCheats)
	/*
		for _, pc := range possibleCheats {
			cg := grid.Copy()
			cg.Set(pc.start, '1')
			cg.Set(pc.end, '2')
			h.PrintGrid(cg)
			fmt.Println(pc)
		}*/
	for _, pc := range possibleCheats {
		if pc.diff >= 100 {
			sum++
		}
	}
	fmt.Println(sum)
}

func part2() {
	sum := 0
	lines := h.GetLinesAsSlice()
	grid := h.ConvertLinesToGrid(lines)
	start := h.Point{}
	end := h.Point{}
	nei := h.NewNeighbourMap[h.Point]()
	grid.ForEachPoint(func(p h.Point) {
		if grid.At(p) == 'S' {
			start = p
		} else if grid.At(p) == 'E' {
			end = p
		}
		if grid.At(p) != '#' {
			for _, n := range p.BasicNeighbours() {
				if n.IsInGrid(grid) && grid.At(n) != '#' {
					nei.AddEdge(p, h.NewEdge(n, 1))
				}
			}
		}
	})
	path, l, distances := Dijkstra(nei, start, end)
	fmt.Println("path length", l)
	for k, v := range distances {
		distances[k] = l - v
	}
	possibleCheats := []Cheat{}
	for i, p := range path {
		fmt.Println(i, "/", len(path))
		for _, pg := range path[i:] {
			md := p.ManhattanDist(pg)
			if md > 20 {
				continue
			}
			if distances[p] <= distances[pg]+md {
				continue
			}
			diff := distances[p] - distances[pg] - md
			if diff < 100 {
				continue
			}
			cheat := Cheat{p, pg, diff}
			possibleCheats = append(possibleCheats, cheat)
		}
	}
	/*
		fmt.Println(possibleCheats)
		for _, r := range h.R(50, 100) {
			count := 0
			for _, pc := range possibleCheats {
				if pc.diff == r {
					count++
				}
			}
			fmt.Println("cheats with", r, "amount", count)
		}
	*/
	/*
		for _, pc := range possibleCheats {
			cg := grid.Copy()
			cg.Set(pc.start, '1')
			cg.Set(pc.end, '2')
			h.PrintGrid(cg)
			fmt.Println(pc)
		}*/
	for _, pc := range possibleCheats {
		if pc.diff >= 100 {
			sum++
		}
	}
	fmt.Println(sum)
}

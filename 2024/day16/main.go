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

const (
	MOVEFORWARD = iota
	TURNRIGHT
	TURNLEFT
)

func part1() {
	sum := 0
	lines := h.GetLinesAsSlice()
	grid := h.ConvertLinesToGrid(lines)
	start := h.Point{}
	end := h.Point{}
	grid.ForEachPoint(func(p h.Point) {
		if grid.At(p) == 'S' {
			start = p
		} else if grid.At(p) == 'E' {
			end = p
		}
	})
	path, l := dijkstra(grid, start, end)
	fmt.Println(path, l)
	fmt.Println(sum)
}

func dijkstra(grid h.Grid, start h.Point, end h.Point) ([]h.Node, int) {
	open := h.PQ[h.Node]{}
	distances := map[h.Node]int{}
	closed := []h.Node{}
	startNode := h.Node{
		start,
		h.EAST,
	}
	open.Push(startNode, 0)
	distances[startNode] = 0
	for open.Len() > 0 {
		next := open.Pop()
		closed = append(closed, next)

		// found end?
		if next.Pos == end {
			return []h.Node{}, distances[next]
		}

		// forward
		forwardPos := next.Pos.Relative(next.Dir)
		if forwardPos.IsInGrid(grid) && grid.At(forwardPos) != '#' {
			n := h.Node{
				forwardPos,
				next.Dir,
			}
			if !slices.Contains(closed, n) {
				open.Push(n, distances[next]+1)
				distances[n] = distances[next] + 1
			}
		}
		// turns
		if next.Dir == h.NORTH || next.Dir == h.SOUTH {
			n0 := h.Node{
				next.Pos,
				h.WEST,
			}
			n1 := h.Node{
				next.Pos,
				h.EAST,
			}
			if !slices.Contains(closed, n0) {
				open.Push(n0, distances[next]+1000)
				distances[n0] = distances[next] + 1000
			}
			if !slices.Contains(closed, n1) {
				open.Push(n1, distances[next]+1000)
				distances[n1] = distances[next] + 1000
			}
		} else {
			n0 := h.Node{
				next.Pos,
				h.NORTH,
			}
			n1 := h.Node{
				next.Pos,
				h.SOUTH,
			}
			if !slices.Contains(closed, n0) {
				open.Push(n0, distances[next]+1000)
				distances[n0] = distances[next] + 1000
			}
			if !slices.Contains(closed, n1) {
				open.Push(n1, distances[next]+1000)
				distances[n1] = distances[next] + 1000
			}
		}
	}
	return []h.Node{}, 0
}

/*
solution is correct but takes to long, ideas:
- dijktra for dist and then use the distance map in combination with FloydWarshal to calculate if point can be on path, but O(3)
- not deterministic idea: randomize the order in which dijktra pushes points to open Priority Queue to change found paths and the let it run until no new paths are found for some time x
- find more conditions for the dfs solution that cancel recursion, like:
  - point exsists on found path and had a smaller distance
  - walks back on currentPath
  - FloydWarshalMatrix + currentLen > maxLen, is O(3) tho

- I guess just one FloydWarshal and then calculation like mentioned would do the trick but still O(3)
- no idea with better runtime yet...
*/
func part2() {
	sum := 0
	lines := h.GetLinesAsSlice()
	grid := h.ConvertLinesToGrid(lines)
	start := h.Point{}
	end := h.Point{}
	grid.ForEachPoint(func(p h.Point) {
		if grid.At(p) == 'S' {
			start = p
		} else if grid.At(p) == 'E' {
			end = p
		}
	})
	_, maxLen := dijkstra(grid, start, end)
	dfsNrOfPathsMaxLen(grid, start, end, start, h.EAST, 0, maxLen, []h.Point{})

	grid.ForEachPoint(func(p h.Point) {
		if grid.At(p) == 'O' {
			sum++
		}
	})
	fmt.Println("part2", sum)
}

func dfsNrOfPathsMaxLen(grid h.Grid, start h.Point, end h.Point, pos h.Point, dir int, l int, maxL int, path []h.Point) int {
	// break
	if l > maxL {
		return 0
	}
	if l == maxL && pos == end {
		fmt.Println("found Path")
		for _, p := range path {
			grid.Set(p, 'O')
		}
		return 1
	}
	// forward
	sumPathsFound := 0
	forwardPos := pos.Relative(dir)
	if forwardPos.IsInGrid(grid) && grid.At(forwardPos) != '#' {
		sumPathsFound += dfsNrOfPathsMaxLen(grid, start, end, forwardPos, dir, l+1, maxL, append(slices.Clone(path), forwardPos))
	}
	// turns
	if dir == h.NORTH || dir == h.SOUTH {
		sumPathsFound += dfsNrOfPathsMaxLen(grid, start, end, pos, h.WEST, l+1000, maxL, append(slices.Clone(path), pos))
		sumPathsFound += dfsNrOfPathsMaxLen(grid, start, end, pos, h.EAST, l+1000, maxL, append(slices.Clone(path), pos))
	} else {
		sumPathsFound += dfsNrOfPathsMaxLen(grid, start, end, pos, h.NORTH, l+1000, maxL, append(slices.Clone(path), pos))
		sumPathsFound += dfsNrOfPathsMaxLen(grid, start, end, pos, h.SOUTH, l+1000, maxL, append(slices.Clone(path), pos))
	}
	return sumPathsFound
}

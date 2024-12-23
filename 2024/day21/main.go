package main

import (
	"AOC/h"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Start")
	part1()
	part2()
}

func getDirOfWalk(from h.Point, to h.Point) int {
	for _, dir := range h.GetBasicDirs() {
		if from.Relative(dir) == to {
			return dir
		}
	}
	return -1
}

func translateBack(char byte) byte {
	switch char {
	case 'n':
		return '^'
	case 's':
		return 'v'
	case 'w':
		return '<'
	case 'e':
		return '>'
	case 'a':
		return 'A'
	}
	return '?'
}

func getCharFromDir(dir int) byte {
	switch dir {
	case h.NORTH:
		return 'n'
	case h.SOUTH:
		return 's'
	case h.WEST:
		return 'w'
	case h.EAST:
		return 'e'
	}
	return ' '
}

func findPositionInGrid(grid h.Grid, c byte) h.Point {
	pf := h.Point{-1, -1}
	grid.ForEachPoint(func(p h.Point) {
		if grid.At(p) == c {
			pf = p
		}
	})
	return pf
}

func genNeis(grid h.Grid) h.NeighbourMap[h.Point] {
	neis := h.NeighbourMap[h.Point]{}
	grid.ForEachPoint(func(p h.Point) {
		if grid.At(p) == ' ' {
			return
		}
		for _, n := range p.BasicNeighbours() {
			if !n.IsInGrid(grid) || grid.At(n) == ' ' {
				continue
			}
			neis.AddEdge(p, h.NewEdge(n, 1))
		}
	})
	return neis
}

var numPadStr = "789\n456\n123\n 0A"
var dirPadStr = " na\nwse"
var numPadGrid = h.ConvertLinesToGrid(strings.Split(numPadStr, "\n"))
var dirPadGrid = h.ConvertLinesToGrid(strings.Split(dirPadStr, "\n"))
var numPadStartPos = h.Point{2, 3}
var dirPadStartPos = h.Point{2, 0}
var dirPadPositions = []h.Point{dirPadStartPos, dirPadStartPos, dirPadStartPos}
var numPadPosition = numPadStartPos

func BFS(nei h.Grid, st h.Point, goal h.Point, doNotEnter h.Point) []h.Point {
	pos := st
	path := []h.Point{pos}
	failed := false
	// x first
	for goal.X > pos.X {
		pos.X++
		path = append(path, pos)
		if pos == doNotEnter {
			failed = true
		}
	}
	for goal.X < pos.X {
		pos.X--
		path = append(path, pos)
		if pos == doNotEnter {
			failed = true
		}
	}
	for goal.Y > pos.Y {
		pos.Y++
		path = append(path, pos)
		if pos == doNotEnter {
			failed = true
		}
	}
	for goal.Y < pos.Y {
		pos.Y--
		path = append(path, pos)
		if pos == doNotEnter {
			failed = true
		}
	}

	if failed {
		pos = st
		path = []h.Point{pos}
		// y first
		for goal.Y > pos.Y {
			pos.Y++

			path = append(path, pos)
		}
		for goal.Y < pos.Y {
			pos.Y--
			path = append(path, pos)
		}
		for goal.X > pos.X {
			pos.X++
			path = append(path, pos)

		}
		for goal.X < pos.X {
			pos.X--
			path = append(path, pos)
		}
	}

	return path
}

func dirPadPath(from h.Point, to h.Point, dirPadIndex int) []h.Point {
	dir := getDirOfWalk(from, to)
	char := getCharFromDir(dir)
	goal := findPositionInGrid(dirPadGrid, char)
	pathDirPad1 := BFS(numPadGrid, dirPadPositions[dirPadIndex], goal, h.Point{0, 0})
	dirPadPositions[dirPadIndex] = goal
	return pathDirPad1
}

func printPath(paths [][]h.Point) {
	fmt.Println("Path: ")
	for _, path := range paths {
		for i := 1; i < len(path); i++ {
			dir := getDirOfWalk(path[i-1], path[i])
			char := getCharFromDir(dir)
			fmt.Print(string(translateBack(char)))
		}
		fmt.Print("A")
	}
	fmt.Println()
}

func part1() {
	sum := 0
	lines := h.GetLinesAsSlice()

	charsToPress := []byte{}

	for _, l := range lines {
		for _, r := range l {
			goal := findPositionInGrid(numPadGrid, byte(r))
			pathToRune := BFS(numPadGrid, numPadPosition, goal, h.Point{0, 3})
			numPadPosition = goal

			pathOnDirPad1 := [][]h.Point{}
			for i := 1; i < len(pathToRune); i++ {
				pathOnDirpadToNextSign := dirPadPath(pathToRune[i-1], pathToRune[i], 0)
				pathOnDirPad1 = append(pathOnDirPad1, pathOnDirpadToNextSign)
			}
			// append back to a
			dijkstraToAPath := BFS(dirPadGrid, dirPadPositions[0], dirPadStartPos, h.Point{0, 0})
			pathOnDirPad1 = append(pathOnDirPad1, dijkstraToAPath)
			dirPadPositions[0] = dirPadStartPos

			pathOnDirPad2 := [][]h.Point{}
			for _, path := range pathOnDirPad1 {
				for i := 1; i < len(path); i++ {
					pathOnDirpadToNextSign := dirPadPath(path[i-1], path[i], 1)
					pathOnDirPad2 = append(pathOnDirPad2, pathOnDirpadToNextSign)
				}
				// append back to a
				dijkstraToAPath = BFS(dirPadGrid, dirPadPositions[1], dirPadStartPos, h.Point{0, 0})
				pathOnDirPad2 = append(pathOnDirPad2, dijkstraToAPath)
				dirPadPositions[1] = dirPadStartPos
			}

			for _, path := range pathOnDirPad2 {
				for i := 1; i < len(path); i++ {
					dir := getDirOfWalk(path[i-1], path[i])
					char := getCharFromDir(dir)
					charsToPress = append(charsToPress, translateBack(char))
				}
				charsToPress = append(charsToPress, 'A')
			}
		}
		seq := strings.Join(h.Map(charsToPress, func(c byte) string { return string(c) }), "")
		fmt.Println(seq)
		n, _ := strconv.Atoi(strings.Split(l, "A")[0])
		sum += len(seq) * n
		charsToPress = []byte{}
	}

	fmt.Println(sum)
}

func part2() {

}

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

func part1() {
	sum := 0
	lines := h.GetLinesAsSlice()
	numPadStr := "789\n456\n123\n 0A"
	dirPadStr := " na\nwse"
	numPadGrid := h.ConvertLinesToGrid(strings.Split(numPadStr, "\n"))
	dirPadGrid := h.ConvertLinesToGrid(strings.Split(dirPadStr, "\n"))
	h.PrintGrid(numPadGrid)
	h.PrintGrid(dirPadGrid)
	numPadNeis := genNeis(numPadGrid)
	dirPadNeis := genNeis(dirPadGrid)
	numPadStartPos := h.Point{2, 3}
	dirPadStartPos := h.Point{2, 0}
	dirPadPositions := []h.Point{dirPadStartPos, dirPadStartPos}
	numPadPosition := numPadStartPos

	charsToPress := []byte{}

	for _, l := range lines {
		for _, r := range l {
			pathNumPad, _ := h.Dijkstra(numPadNeis, numPadPosition, findPositionInGrid(numPadGrid, byte(r)))
			numPadPosition = numPadStartPos
			for i, newP := range pathNumPad {
				if i == 0 {
					continue
				}
				dir := getDirOfWalk(pathNumPad[i-1], newP)
				char := getCharFromDir(dir)
				pathDirPad1, _ := h.Dijkstra(dirPadNeis, dirPadPositions[0], findPositionInGrid(dirPadGrid, char))
				dirPadPositions[0] = pathDirPad1[len(pathDirPad1)-1]
				toA, _ := h.Dijkstra(dirPadNeis, dirPadPositions[0], dirPadStartPos)
				pathDirPad1 = append(pathDirPad1, toA[1:]...)
				dirPadPositions[0] = dirPadStartPos
				for j, newP1 := range pathDirPad1 {
					if j == 0 {
						continue
					}
					dir1 := getDirOfWalk(pathDirPad1[j-1], newP1)
					char1 := getCharFromDir(dir1)
					pathDirPad2, _ := h.Dijkstra(dirPadNeis, dirPadPositions[1], findPositionInGrid(dirPadGrid, char1))
					dirPadPositions[1] = pathDirPad2[len(pathDirPad2)-1]
					pressPosition := dirPadPositions[1]
					toA1, _ := h.Dijkstra(dirPadNeis, dirPadPositions[1], dirPadStartPos)
					pathDirPad2 = append(pathDirPad2, toA1[1:]...)
					dirPadPositions[1] = dirPadStartPos
					for k, newP2 := range pathDirPad2 {
						if k == 0 {
							continue
						}
						if newP2 == pressPosition {
							fmt.Print(string('A'))
						}
						dir2 := getDirOfWalk(pathDirPad2[k-1], newP2)
						char2 := getCharFromDir(dir2)
						fmt.Print(string(translateBack(char2)))
						charsToPress = append(charsToPress, char2)
					}
					fmt.Print(string('A'))
				}
			}
		}
	}

	fmt.Println(charsToPress)
	fmt.Println(sum)
}

func part2() {

}

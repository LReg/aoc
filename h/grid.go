package h

import (
	"fmt"
	"strconv"
)

type Grid [][]byte

// GridCompareByteArr this is good for easy comparing of byte array in a grid
func (grid Grid) GridCompareByteArr(p Point, dir int, val []byte) bool {
	same := false
	WalkThroughLine(grid, dir, p, func(pos Point, i int) bool {
		if grid[pos.X][pos.Y] != val[i] {
			return false
		}
		if i == len(val)-1 {
			same = true
			return false
		}
		return true
	})
	return same
}

func (grid Grid) GridCompareStr(p Point, dir int, val string) bool {
	return grid.GridCompareByteArr(p, dir, []byte(val))
}

func (grid Grid) HashValue() int {
	hash := 0
	for i, row := range grid {
		for j, b := range row {
			hash += i + j + int(b)
		}
	}
	return hash
}

func IsPointInGrid[T any](grid [][]T, p Point) bool {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return false
	}
	if p.Y < 0 || p.Y >= len(grid[0]) {
		return false
	}
	if p.X < 0 || p.X >= len(grid) {
		return false
	}
	return true
}

func (grid Grid) ForEachPoint(f func(p Point)) {
	WalkThrough(grid, EAST, Point{0, 0}, f)
}

func (grid Grid) At(p Point) byte {
	return grid[p.X][p.Y]
}

func (grid Grid) AtNum(p Point) int {
	n, _ := strconv.Atoi(string(grid.At(p)))
	return n
}

// Set this mutates!
func (grid Grid) Set(p Point, b byte) {
	grid[p.X][p.Y] = b
}

func (grid Grid) Neighbours(p Point) []byte {
	n := make([]byte, 0)
	for _, dir := range GetAllDirs() {
		p := p.Relative(dir)
		if IsPointInGrid(grid, p) {
			n = append(n, grid.At(p))
		}
	}
	return n
}

func (grid Grid) BasicNeighbours(p Point) []byte {
	n := make([]byte, 0)
	for _, dir := range GetBasicDirs() {
		n = append(n, grid.At(p.Relative(dir)))
	}
	return n
}

func WalkThrough[T any](grid [][]T, dir int, start Point, f func(p Point)) {
	if dir == EAST {
		for y := start.Y; y < len(grid[0]); y++ {
			for x := start.X; x < len(grid); x++ {
				f(Point{x, y})
			}
		}
	} else if dir == NORTH {
		for x := start.X; x < len(grid); x++ {
			for y := start.Y; y < len(grid[0]); y++ {
				f(Point{x, y})
			}
		}
	} else if dir == NORTHEAST {
		fmt.Println("[Warning]: Diagonal not implemented")
	}
}

// WalkThroughLine bool at function return true -> continue, false -> break
// looking back this function seems rather senseless for usage in scripts
// it is rather just useful for the GridCompareByteArr function to make it maintainable
func WalkThroughLine[T any](grid [][]T, dir int, start Point, f func(p Point, i int) bool) {
	i := 0
	if !IsPointInGrid(grid, start) {
		return
	}
	if dir == EAST {
		for x := start.X; x < len(grid); x++ {
			b := f(Point{x, start.Y}, i)
			i++
			if !b {
				return
			}
		}
	} else if dir == WEST {
		for x := start.X; x >= 0; x-- {
			b := f(Point{x, start.Y}, i)
			i++
			if !b {
				return
			}
		}
	} else if dir == NORTH {
		for y := start.Y; y < len(grid[0]); y++ {
			b := f(Point{start.X, y}, i)
			i++
			if !b {
				return
			}
		}
	} else if dir == SOUTH {
		for y := start.Y; y >= 0; y-- {
			b := f(Point{start.X, y}, i)
			i++
			if !b {
				return
			}
		}
	} else if dir == NORTHEAST {
		p := Point{start.X, start.Y}
		for IsPointInGrid(grid, p) {
			b := f(p, i)
			i++
			p.X++
			p.Y++
			if !b {
				return
			}
		}
	} else if dir == SOUTHEAST {
		p := Point{start.X, start.Y}
		for IsPointInGrid(grid, p) {
			b := f(p, i)
			i++
			p.X++
			p.Y--
			if !b {
				return
			}
		}
	} else if dir == SOUTHWEST {
		p := Point{start.X, start.Y}
		for IsPointInGrid(grid, p) {
			b := f(p, i)
			i++
			p.X--
			p.Y--
			if !b {
				return
			}
		}
	} else if dir == NORTHWEST {
		p := Point{start.X, start.Y}
		for IsPointInGrid(grid, p) {
			b := f(p, i)
			i++
			p.X--
			p.Y++
			if !b {
				return
			}
		}
	}
}

func Create3DIntGrid() [][][]int {
	slice := make([][][]int, 0)
	slice = append(slice, make([][]int, 0))
	slice[0] = append(slice[0], make([]int, 0))
	return slice
}

func Create3DByteGrid() [][][]byte {
	slice := make([][][]byte, 0)
	slice = append(slice, make([][]byte, 0))
	slice[0] = append(slice[0], make([]byte, 0))
	return slice
}

package h

import "fmt"

const (
	HORIZONTAL = iota
	VERTICAL
	DIAGONALUP
	DIAGONALDOWN
	HORIZONTALREVERSE
	VERTICALREVERSE
	DIAGONALUPREVERSE
	DIAGONALDOWNREVERSE
)

// GridCompareByte this is good for easy comparing of byte array in a grid
func GridCompareByte(grid [][]byte, p Point, dir int, val []byte) bool {
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

func WalkThrough[T any](grid [][]T, dir int, start Point, f func(p Point)) {
	if dir == HORIZONTAL {
		for y := start.Y; y < len(grid[0]); y++ {
			for x := start.X; x < len(grid); x++ {
				f(Point{x, y})
			}
		}
	} else if dir == VERTICAL {
		for x := start.X; x < len(grid); x++ {
			for y := start.Y; y < len(grid[0]); y++ {
				f(Point{x, y})
			}
		}
	} else if dir == DIAGONALUP {
		fmt.Println("[Warning]: Diagonal not implemented")
	}
}

// WalkThroughLine bool at function return true -> continue, false -> break
// looking back this function seems rather senseless for usage in scripts
// it is rather just useful for the GridCompareByte function to make it maintainable
func WalkThroughLine[T any](grid [][]T, dir int, start Point, f func(p Point, i int) bool) {
	i := 0
	if !IsPointInGrid(grid, start) {
		return
	}
	if dir == HORIZONTAL {
		for x := start.X; x < len(grid); x++ {
			b := f(Point{x, start.Y}, i)
			i++
			if !b {
				return
			}
		}
	} else if dir == HORIZONTALREVERSE {
		for x := start.X; x >= 0; x-- {
			b := f(Point{x, start.Y}, i)
			i++
			if !b {
				return
			}
		}
	} else if dir == VERTICAL {
		for y := start.Y; y < len(grid[0]); y++ {
			b := f(Point{start.X, y}, i)
			i++
			if !b {
				return
			}
		}
	} else if dir == VERTICALREVERSE {
		for y := start.Y; y >= 0; y-- {
			b := f(Point{start.X, y}, i)
			i++
			if !b {
				return
			}
		}
	} else if dir == DIAGONALUP {
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
	} else if dir == DIAGONALDOWN {
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
	} else if dir == DIAGONALDOWNREVERSE {
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
	} else if dir == DIAGONALUPREVERSE {
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

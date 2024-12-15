package h

import (
	"cmp"
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

func (grid Grid) Copy() Grid {
	newGrid := make([][]byte, len(grid))
	for i, row := range grid {
		newGrid[i] = make([]byte, len(row))
		copy(newGrid[i], row)
	}
	return newGrid
}

// return count changed
func (grid Grid) FloodFillBasic(start Point, b byte, stopHere func(p Point) bool) int {
	if !IsPointInGrid(grid, start) {
		return 0
	}
	if grid[start.X][start.Y] == b {
		return 0
	}
	if stopHere(start) {
		return 0
	}

	grid[start.X][start.Y] = b
	sumChanged := 1

	for _, dir := range GetBasicDirs() {
		sumChanged += grid.FloodFillBasic(start.Relative(dir), b, stopHere)
	}
	return sumChanged
}

func (grid Grid) DijkstraPosNum(start Point, end Point) ([]Point, int) {
	return grid.Dijkstra(start, end, func(p Point) int {
		return grid.AtNum(p)
	})
}

func (grid Grid) ProduceNeighbourMap(weight func(p Point) int) map[Point][]Edge[Point] {
	neighbourMap := make(map[Point][]Edge[Point])
	grid.ForEachPoint(func(p Point) {
		neighbourMap[p] = make([]Edge[Point], 0)
		for _, dir := range GetBasicDirs() {
			neighbour := p.Relative(dir)
			if IsPointInGrid(grid, neighbour) {
				weight := weight(neighbour)
				neighbourMap[p] = append(neighbourMap[p], Edge[Point]{neighbour, weight})
			}
		}
	})
	return neighbourMap
}

func (grid Grid) FloydWarshallPosNum() FWMatrix[Point] {
	return FloydWarshall(grid.ProduceNeighbourMap(func(p Point) int {
		return grid.AtNum(p)
	}))
}

func (grid Grid) FloydWarshall(weight func(p Point) int) FWMatrix[Point] {
	return FloydWarshall(grid.ProduceNeighbourMap(weight))
}

func (grid Grid) Dijkstra(start Point, end Point, weight func(p Point) int) ([]Point, int) {
	return DijkstraOld[Point](grid.ProduceNeighbourMap(weight), start, end)
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
		for y := start.Y; y < len(grid[0]); y-- {
			b := f(Point{start.X, y}, i)
			i++
			if !b {
				return
			}
		}
	} else if dir == SOUTH {
		for y := start.Y; y >= 0; y++ {
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
			p.Y--
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
			p.Y++
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
			p.Y++
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
			p.Y--
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

func CreateGrid(xMax int, yMax int) Grid {
	grid := make([][]byte, xMax)
	for i, _ := range grid {
		grid[i] = make([]byte, yMax)
	}
	return grid
}

func CreateOrderedGrid[T cmp.Ordered](xMax int, yMax int) [][]T {
	grid := make([][]T, xMax)
	for i, _ := range grid {
		grid[i] = make([]T, yMax)
	}
	return grid
}

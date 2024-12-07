package h

import (
	"math"
)

// point will not mutate

type Point struct {
	X, Y int
}

type Point3 struct {
	X, Y, Z int
}

func (p Point) RelativeN(dir int, n int) Point {
	switch dir {
	case NORTH:
		p.Y -= n
		break
	case NORTHEAST:
		p.X += n
		p.Y -= n
		break
	case EAST:
		p.X += n
		break
	case SOUTHEAST:
		p.X += n
		p.Y += n
		break
	case SOUTH:
		p.Y += n
		break
	case SOUTHWEST:
		p.X -= n
		p.Y += n
		break
	case WEST:
		p.X -= n
		break
	case NORTHWEST:
		p.X -= n
		p.Y -= n
		break
	}
	return p
}

func (p Point) Neighbours() []Point {
	res := make([]Point, 0)
	for _, dir := range GetAllDirs() {
		n := p.Relative(dir)
		res = append(res, n)
	}
	return res
}

func (p Point) BasicNeighbours() []Point {
	res := make([]Point, 0)
	for _, dir := range GetBasicDirs() {
		n := p.Relative(dir)
		res = append(res, n)
	}
	return res
}

func (p Point) IsInGrid(grid Grid) bool {
	return IsPointInGrid(grid, p)
}

func (p Point) Relative(dir int) Point {
	return p.RelativeN(dir, 1)
}

func (p Point) distSqrt(y Point) int {
	dx := p.X - y.X
	dy := p.Y - y.Y
	return dx*dx + dy*dy
}

func (p Point3) distSqrt(y Point3) int {
	dx := p.X - y.X
	dy := p.Y - y.Y
	dz := p.Z - y.Z
	return dx*dx + dy*dy + dz*dz
}

func (p Point) Dist(y Point) float64 {
	return math.Sqrt(float64(p.distSqrt(y)))
}

func (p Point3) Dist(y Point3) float64 {
	return math.Sqrt(float64(p.distSqrt(y)))
}

func (p Point) ManhattanDist(y Point) int {
	return Abs(p.X-y.X) + Abs(p.Y-y.Y)
}

func (p Point3) ManhattanDist(y Point3) int {
	return Abs(p.X-y.X) + Abs(p.Y-y.Y) + Abs(p.Z-y.Z)
}

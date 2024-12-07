package h

import (
	"math"
	"slices"
)

type Edge struct {
	To     Point
	Weight int
}

func Dijkstra(neighbourMap map[Point][]Edge, start Point, end Point) ([]Point, int) {
	distances := make(map[Point]int)
	for point := range neighbourMap {
		distances[point] = math.MaxInt
	}
	distances[start] = 0
	previos := make(map[Point]Point)

	unvisited := make([]Point, 1)
	unvisited[0] = start

	for len(unvisited) > 0 {
		// find the node with the smallest distance
		smallesDistance := math.MaxInt
		nearestPoint := Point{}
		for point, distance := range distances {
			if smallesDistance > distance {
				smallesDistance = distance
				nearestPoint = point
			}
		}

		if nearestPoint == end {
			break
		}

		// remove the node from the unvisited list
		unvisited = slices.DeleteFunc(unvisited, func(p Point) bool {
			return p == nearestPoint
		})

		for _, neighbour := range neighbourMap[nearestPoint] {
			newDistance := distances[nearestPoint] + neighbour.Weight
			if newDistance < distances[neighbour.To] {
				distances[neighbour.To] = newDistance
				previos[neighbour.To] = nearestPoint
			}
		}

	}

	path := make([]Point, 0)
	for point := end; point != start; point = previos[point] {
		path = append(path, point)
	}

	path = append(path, start)

	return path, distances[end]
}

package h

import (
	"cmp"
	"math"
	"slices"
)

type Edge[T cmp.Ordered] struct {
	To     Point
	Weight int
}

func Dijkstra[T cmp.Ordered](neighbourMap map[Point][]Edge[T], start Point, end Point, breakWhenFound bool) ([]Point, int) {
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
		nearestPoint := unvisited[0]

		if nearestPoint == end {
			break
		}

		// remove the node from the unvisited list
		unvisited = slices.DeleteFunc(unvisited, func(p Point) bool {
			return p == nearestPoint
		})

		// update neighbour distances and add them to the unvisited list
		for _, neighbour := range neighbourMap[nearestPoint] {
			newDistance := distances[nearestPoint] + neighbour.Weight
			if newDistance < distances[neighbour.To] {
				distances[neighbour.To] = newDistance
				previos[neighbour.To] = nearestPoint
				if !slices.Contains(unvisited, neighbour.To) {
					if len(unvisited) == 0 {
						unvisited = append(unvisited, neighbour.To)
					} else {
						for i, un := range unvisited {
							distOfUnvisitedInArr := distances[un]
							if distances[neighbour.To] <= distOfUnvisitedInArr {
								if i-1 < 0 {
									unvisited = slices.Insert(unvisited, 0, neighbour.To)
									break
								} else {
									unvisited = slices.Insert(unvisited, i-1, neighbour.To)
									break
								}
							}
						}
					}
				}
			}
			if !slices.Contains(unvisited, neighbour.To) {
				unvisited = append(unvisited, neighbour.To)
			}
		}
	}

	path := make([]Point, 0)
	for point := end; point != start; point = previos[point] {
		path = append(path, point)
	}
	path = append(path, start)
	slices.Reverse(path)

	return path, distances[end]
}

// DijkstraOld runs better did not manage to improve the new one
func DijkstraOld[T cmp.Ordered](neighbourMap map[Point][]Edge[T], start Point, end Point, breakWhenFound bool) ([]Point, int) {
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
		nearestPoint := Point{}
		minDist := math.MaxInt
		for _, un := range unvisited {
			if distances[un] < minDist {
				minDist = distances[un]
				nearestPoint = un
			}
		}

		if nearestPoint == end {
			break
		}

		// remove the node from the unvisited list
		unvisited = slices.DeleteFunc(unvisited, func(p Point) bool {
			return p == nearestPoint
		})

		// update neighbour distances and add them to the unvisited list
		for _, neighbour := range neighbourMap[nearestPoint] {
			newDistance := distances[nearestPoint] + neighbour.Weight
			if newDistance < distances[neighbour.To] {
				distances[neighbour.To] = newDistance
				previos[neighbour.To] = nearestPoint
				if !slices.Contains(unvisited, neighbour.To) {
					unvisited = append(unvisited, neighbour.To)
				}
			}
		}
	}

	path := make([]Point, 0)
	for point := end; point != start; point = previos[point] {
		path = append(path, point)
	}
	path = append(path, start)
	slices.Reverse(path)

	return path, distances[end]
}

func CrossProduct[T any](possibleElements []T, n int) [][]T {
	possibleCombinations := make([][]T, 1)
	possibleCombinations[0] = []T{}
	for i := 0; i < n; i++ {
		refreshedPossibleCombinations := make([][]T, 0)
		for _, possibleCombination := range possibleCombinations {
			for _, possibleElement := range possibleElements {
				newComb := append([]T{}, possibleCombination...)
				newComb = append(newComb, possibleElement)
				refreshedPossibleCombinations = append(refreshedPossibleCombinations, newComb)
			}
		}
		possibleCombinations = refreshedPossibleCombinations
	}
	return possibleCombinations
}

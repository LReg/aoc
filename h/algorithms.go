package h

import (
	"cmp"
	"fmt"
	"math"
	"slices"
	"strconv"
)

type Edge[T cmp.Ordered | Point] struct {
	To     T
	Weight int
}
type NeighbourMap[T cmp.Ordered | Point] map[T][]Edge[T]

func Dijkstra[T cmp.Ordered | Point](neighbourMap map[T][]Edge[T], start T, end T, breakWhenFound bool) ([]T, int) {
	distances := make(map[T]int)
	for point := range neighbourMap {
		distances[point] = math.MaxInt
	}
	distances[start] = 0
	previos := make(map[T]T)

	unvisited := make([]T, 1)
	unvisited[0] = start

	for len(unvisited) > 0 {
		// find the node with the smallest distance
		nearestPoint := unvisited[0]

		if nearestPoint == end {
			break
		}

		// remove the node from the unvisited list
		unvisited = slices.DeleteFunc(unvisited, func(p T) bool {
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

	path := make([]T, 0)
	for point := end; point != start; point = previos[point] {
		path = append(path, point)
	}
	path = append(path, start)
	slices.Reverse(path)

	return path, distances[end]
}

// DijkstraOld runs better did not manage to improve the new one
func DijkstraOld[T cmp.Ordered | Point](neighbourMap map[T][]Edge[T], start T, end T, breakWhenFound bool) ([]T, int) {
	distances := make(map[T]int)
	for point := range neighbourMap {
		distances[point] = math.MaxInt
	}
	distances[start] = 0
	previos := make(map[T]T)

	unvisited := make([]T, 1)
	unvisited[0] = start

	for len(unvisited) > 0 {
		// find the node with the smallest distance
		var nearestPoint T
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
		unvisited = slices.DeleteFunc(unvisited, func(p T) bool {
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

	path := make([]T, 0)
	for point := end; point != start; point = previos[point] {
		path = append(path, point)
	}
	path = append(path, start)
	slices.Reverse(path)

	return path, distances[end]
}

type FWMatrix[T cmp.Ordered | Point] map[T]map[T]int

func (m FWMatrix[string]) PrintStr() {

	fmt.Println("-------")

	allVertex := make([]string, 0)
	allVertexReversed := make([]string, 0)
	for to := range m {
		allVertex = append(allVertex, to)
		allVertexReversed = append(allVertexReversed, to)
	}
	slices.Reverse(allVertexReversed)

	fmt.Print("Matrix")
	fmt.Print("\t\t\t")
	for _, from := range allVertex {
		fmt.Print(from)
		fmt.Print("\t")
	}
	fmt.Println()
	for _, to := range allVertex {
		fmt.Print(to)
		fmt.Print("\t\t\t")
		for _, from := range allVertex {
			if m[from][to] == math.MaxInt {
				fmt.Print("MAX" + "\t")
			} else {
				fmt.Print(strconv.Itoa(m[from][to]) + "\t")
			}
		}
		fmt.Println()
	}

	fmt.Println("-------")
}

func FloydWarshall[T cmp.Ordered | Point](neighbourMap NeighbourMap[T]) FWMatrix[T] {
	matrix := make(FWMatrix[T])
	allVertex := make([]T, 0)
	for from, edges := range neighbourMap {
		if !slices.Contains(allVertex, from) {
			allVertex = append(allVertex, from)
		}
		for _, e := range edges {
			if !slices.Contains(allVertex, e.To) {
				allVertex = append(allVertex, e.To)
			}
		}
	}

	for _, from := range allVertex {
		matrix[from] = make(map[T]int)
		for _, to := range allVertex {
			if from == to {
				matrix[from][to] = 0
			} else {
				matrix[from][to] = math.MaxInt
			}
		}
	}

	for from, edges := range neighbourMap {
		for _, edge := range edges {
			matrix[from][edge.To] = edge.Weight
		}
	}

	i := 0
	for k, _ := range neighbourMap {
		fmt.Println(i, "/", len(neighbourMap))
		i++
		for i, _ := range neighbourMap {
			for j, _ := range neighbourMap {
				n := matrix[i][k] + matrix[k][j]
				if n < 0 {
					continue
				}
				matrix[i][j] = Min(matrix[i][j], n)
			}
		}
	}
	return matrix
}

func Permutations[T any](possibleElements []T, n int) [][]T {
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

func TSPReturnToStart[T cmp.Ordered | Point](neighbourMap NeighbourMap[T]) ([]T, int) {

	allVertex := make([]T, 0)
	for from, edges := range neighbourMap {
		if !slices.Contains(allVertex, from) {
			allVertex = append(allVertex, from)
		}
		for _, e := range edges {
			if !slices.Contains(allVertex, e.To) {
				allVertex = append(allVertex, e.To)
			}
		}
	}

	bestPath := make([]T, 0)
	bestWeight := math.MaxInt

	// TODO

	return bestPath, bestWeight
}

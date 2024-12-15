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

func NewEdge[T cmp.Ordered | Point](to T, weight int) Edge[T] {
	return Edge[T]{To: to, Weight: weight}
}

type NeighbourMap[T cmp.Ordered | Point] map[T][]Edge[T]

func NewNeighbourMap[T cmp.Ordered | Point]() NeighbourMap[T] {
	return make(NeighbourMap[T])
}
func (n NeighbourMap[T]) AddEdge(from T, e Edge[T]) NeighbourMap[T] {
	_, ok := n[from]
	if !ok {
		n[from] = []Edge[T]{}
	}
	n[from] = append(n[from], e)
	return n
}
func (n NeighbourMap[T]) CanTravel(from T, to T) bool {
	for _, next := range n[from] {
		if next.To == to {
			return true
		}
	}
	return false
}

func (n NeighbourMap[T]) Weight(from T, to T) int {
	for _, e := range n[from] {
		if e.To == to {
			return e.Weight
		}
	}
	return -1
}

func (n NeighbourMap[T]) AllVertex() []T {
	allVertex := make([]T, 0)
	for from, edges := range n {
		if !slices.Contains(allVertex, from) {
			allVertex = append(allVertex, from)
		}
		for _, e := range edges {
			if !slices.Contains(allVertex, e.To) {
				allVertex = append(allVertex, e.To)
			}
		}
	}
	return allVertex
}

func Dijkstra[T cmp.Ordered | Point](neighbourMap map[T][]Edge[T], start T, end T) ([]T, int) {
	distances := make(map[T]int)
	for point := range neighbourMap {
		distances[point] = math.MaxInt
	}
	distances[start] = 0
	previos := make(map[T]T)

	unvisited := NewPC[T]()
	unvisited.Push(start, 0)
	visited := make(map[T]bool)

	for unvisited.Len() > 0 {
		nearestPoint := unvisited.Pop()

		if nearestPoint == end {
			break
		}

		visited[nearestPoint] = true

		// update neighbour distances and add them to the unvisited list
		for _, neighbour := range neighbourMap[nearestPoint] {
			if visited[neighbour.To] {
				continue
			}
			newDistance := distances[nearestPoint] + neighbour.Weight
			if newDistance < distances[neighbour.To] {
				distances[neighbour.To] = newDistance
				previos[neighbour.To] = nearestPoint
				if !unvisited.Contains(neighbour.To) {
					unvisited.Push(neighbour.To, newDistance)
				} else {
					unvisited.UpdatePriority(neighbour.To, newDistance)
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

// DijkstraOld runs better did not manage to improve the new one
func DijkstraOld[T cmp.Ordered | Point](neighbourMap map[T][]Edge[T], start T, end T) ([]T, int) {
	distances := make(map[T]int)
	for point := range neighbourMap {
		distances[point] = math.MaxInt
	}
	distances[start] = 0
	previos := make(map[T]T)

	unvisited := make([]T, 1)
	unvisited[0] = start
	visited := make(map[T]bool)

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

		visited[nearestPoint] = true

		// remove the node from the unvisited list
		unvisited = slices.DeleteFunc(unvisited, func(p T) bool {
			return p == nearestPoint
		})

		// update neighbour distances and add them to the unvisited list
		for _, neighbour := range neighbourMap[nearestPoint] {
			if visited[neighbour.To] {
				continue
			}
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

func (m FWMatrix[T]) PrintStr() {
	fmt.Println("-------")

	allVertex := make([]T, 0)
	allVertexReversed := make([]T, 0)
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
	allVertex := neighbourMap.AllVertex()

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
	allVertex := neighbourMap.AllVertex()

	bestPath := make([]T, 0)
	bestWeight := math.MaxInt

	var tspHelper func(path []T, visited map[T]bool, currentWeight int)
	tspHelper = func(path []T, visited map[T]bool, currentWeight int) {
		if len(path) == len(allVertex) {
			start := path[0]
			last := path[len(path)-1]
			for _, edge := range neighbourMap[last] {
				if edge.To == start {
					totalWeight := currentWeight + edge.Weight
					if totalWeight < bestWeight {
						bestWeight = totalWeight
						bestPath = append([]T{}, path...)
						bestPath = append(bestPath, start)
					}
					break
				}
			}
			return
		}

		current := path[len(path)-1]
		for _, edge := range neighbourMap[current] {
			if !visited[edge.To] {
				visited[edge.To] = true
				tspHelper(append(slices.Clone(path), edge.To), visited, currentWeight+edge.Weight)
				visited[edge.To] = false
			}
		}
	}

	for _, start := range allVertex {
		visited := make(map[T]bool)
		visited[start] = true
		tspHelper([]T{start}, visited, 0)
	}

	return bestPath, bestWeight
}

func TSPReturnToStartLongestPath[T cmp.Ordered | Point](neighbourMap NeighbourMap[T]) ([]T, int) {
	allVertex := neighbourMap.AllVertex()

	bestPath := make([]T, 0)
	bestWeight := 0

	var tspHelper func(path []T, visited map[T]bool, currentWeight int)
	tspHelper = func(path []T, visited map[T]bool, currentWeight int) {
		if len(path) == len(allVertex) {
			start := path[0]
			last := path[len(path)-1]
			for _, edge := range neighbourMap[last] {
				if edge.To == start {
					totalWeight := currentWeight + edge.Weight
					if totalWeight > bestWeight {
						bestWeight = totalWeight
						bestPath = append([]T{}, path...)
						bestPath = append(bestPath, start)
					}
					break
				}
			}
			return
		}

		current := path[len(path)-1]
		for _, edge := range neighbourMap[current] {
			if !visited[edge.To] {
				visited[edge.To] = true
				tspHelper(append(slices.Clone(path), edge.To), visited, currentWeight+edge.Weight)
				visited[edge.To] = false
			}
		}
	}

	for _, start := range allVertex {
		visited := make(map[T]bool)
		visited[start] = true
		tspHelper([]T{start}, visited, 0)
	}

	return bestPath, bestWeight
}

func TSP[T cmp.Ordered | Point](neighbourMap NeighbourMap[T]) ([]T, int) {
	// Alle Knoten sammeln
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

	var tspHelper func(path []T, visited map[T]bool, currentWeight int)
	tspHelper = func(path []T, visited map[T]bool, currentWeight int) {
		if len(path) == len(allVertex) {
			if currentWeight < bestWeight {
				bestWeight = currentWeight
				bestPath = append([]T{}, path...)
			}
			return
		}

		current := path[len(path)-1]
		for _, edge := range neighbourMap[current] {
			if !visited[edge.To] {
				visited[edge.To] = true
				tspHelper(append(slices.Clone(path), edge.To), visited, currentWeight+edge.Weight)
				visited[edge.To] = false
			}
		}
	}

	for _, start := range allVertex {
		visited := make(map[T]bool)
		visited[start] = true
		tspHelper([]T{start}, visited, 0)
	}

	return bestPath, bestWeight
}

func TSPLongestPath[T cmp.Ordered | Point](neighbourMap NeighbourMap[T]) ([]T, int) {
	// Alle Knoten sammeln
	allVertex := neighbourMap.AllVertex()

	bestPath := make([]T, 0)
	bestWeight := 0

	var tspHelper func(path []T, visited map[T]bool, currentWeight int)
	tspHelper = func(path []T, visited map[T]bool, currentWeight int) {
		if len(path) == len(allVertex) {
			if currentWeight > bestWeight {
				bestWeight = currentWeight
				bestPath = append([]T{}, path...)
			}
			return
		}

		current := path[len(path)-1]
		for _, edge := range neighbourMap[current] {
			if !visited[edge.To] {
				visited[edge.To] = true
				tspHelper(append(slices.Clone(path), edge.To), visited, currentWeight+edge.Weight)
				visited[edge.To] = false
			}
		}
	}

	for _, start := range allVertex {
		visited := make(map[T]bool)
		visited[start] = true
		tspHelper([]T{start}, visited, 0)
	}

	return bestPath, bestWeight
}

func DFSNrOfPaths[T cmp.Ordered | Point](nei NeighbourMap[T], st T, goalCondition func(p T) bool, dep int, maxD int) int {
	if dep == maxD {
		return 0
	}
	if goalCondition(st) {
		return 1
	}
	neis := nei[st]
	nFound := 0
	for _, n := range neis {
		nr := DFSNrOfPaths(nei, n.To, goalCondition, dep+1, maxD)
		nFound += nr
	}
	return nFound
}

func DFSAnyFoundPath[T cmp.Ordered | Point](nei NeighbourMap[T], st T, goalCondition func(p T) bool, dep int, maxD int) bool {
	_, l := DFS(nei, st, goalCondition, dep, maxD)
	return l != math.MaxInt
}

func DFS[T cmp.Ordered | Point](nei NeighbourMap[T], st T, goalCondition func(p T) bool, dep int, maxD int) ([]T, int) {
	if dep == maxD {
		return []T{}, math.MaxInt
	}
	if goalCondition(st) {
		return []T{st}, 0
	}
	neis := nei[st]
	for _, n := range neis {
		path, length := DFS(nei, n.To, goalCondition, dep+1, maxD)
		if len(path) > 0 {
			return append([]T{st}, path...), length + 1
		}
	}
	return []T{}, math.MaxInt
}

func BFS[T cmp.Ordered | Point](nei NeighbourMap[T], st T, goalCondition func(p T) bool) ([]T, int) {
	distances := map[T]int{}
	visited := map[T]bool{}
	previous := map[T]T{}
	open := []T{st}

	visited[st] = true
	distances[st] = 0

	for len(open) > 0 {
		el := open[0]
		open = open[1:]
		visited[el] = true

		for _, e := range nei[el] {
			if visited[e.To] || slices.Contains(open, e.To) {
				continue
			}
			distances[e.To] = distances[el] + e.Weight
			previous[e.To] = el

			if goalCondition(e.To) {
				path := make([]T, 0)
				for point := e.To; point != st; point = previous[point] {
					path = append(path, point)
				}
				path = append(path, st)
				slices.Reverse(path)
				return path, distances[e.To]
			}

			open = append(open, e.To)
		}
	}
	return []T{}, math.MaxInt
}

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

func part1() {
	lines := h.GetLinesAsSlice()
	edges := h.NeighbourMap[string]{}
	for _, l := range lines {
		p := strings.Split(l, "=")
		w, _ := strconv.Atoi(strings.TrimSpace(p[1]))
		fromto := strings.Split(p[0], " to ")
		from := strings.TrimSpace(fromto[0])
		to := strings.TrimSpace(fromto[1])
		_, ok := edges[from]
		e := h.Edge[string]{
			To:     to,
			Weight: w,
		}
		if !ok {
			edges[from] = append([]h.Edge[string]{}, e)
		} else {
			edges[from] = append(edges[from], e)
		}
	}
	path, l := h.TSPReturnToStart(edges)
	fmt.Println(path)

	fmt.Println(l)
}

func part2() {

}

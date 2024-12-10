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
	//part1DiffrentApproach()
	part2()
}

func part1() {
	lines := h.GetLinesAsSlice()
	neis := h.NewNeighbourMap[string]()
	for _, l := range lines {
		p := strings.Split(l, "=")
		w, _ := strconv.Atoi(strings.TrimSpace(p[1]))
		fromto := strings.Split(p[0], " to ")
		from := strings.TrimSpace(fromto[0])
		to := strings.TrimSpace(fromto[1])
		neis.AddEdge(from, h.NewEdge(to, w))
		neis.AddEdge(to, h.NewEdge(from, w))
	}
	path, l := h.TSP(neis)
	fmt.Println(path)
	fmt.Println(l)
}

func part2() {
	lines := h.GetLinesAsSlice()
	neis := h.NewNeighbourMap[string]()
	for _, l := range lines {
		p := strings.Split(l, "=")
		w, _ := strconv.Atoi(strings.TrimSpace(p[1]))
		fromto := strings.Split(p[0], " to ")
		from := strings.TrimSpace(fromto[0])
		to := strings.TrimSpace(fromto[1])
		neis.AddEdge(from, h.NewEdge(to, w))
		neis.AddEdge(to, h.NewEdge(from, w))
	}
	path, l := h.TSPLongestPath(neis)
	fmt.Println(path)
	fmt.Println(l)
}

package main

import (
	"AOC/h"
	"fmt"
	"slices"
)

func main() {
	fmt.Println("Start")
	part1()
	part2()
}

func part1() {
	sum := 0
	lines := h.GetLinesAsSlice()
	for _, l := range lines {
		lwh := h.ExtrapolateNumbersFromString(l, "x")
		s := []int{lwh[0] * lwh[1], lwh[1] * lwh[2], lwh[0] * lwh[2]}
		slack := slices.Min(s)
		flaeche := h.Reduce(h.Map(s, func(i int) int { return i * 2 }), h.ReduceSumInt, 0)
		sum += flaeche + slack
	}
	fmt.Println(sum)
}

func part2() {
	sum := 0
	lines := h.GetLinesAsSlice()
	for _, l := range lines {
		lwh := h.ExtrapolateNumbersFromString(l, "x")
		bow := h.Reduce(lwh, func(prev int, curr int) int { return prev * curr }, 1)
		ribbon := 0
		for _, _ = range h.Iter(2) {
			m := slices.Min(lwh)
			ribbon += m + m
			index := slices.Index(lwh, m)
			lwh = h.DeleteIndexFromSlice(lwh, index)
		}
		sum += ribbon + bow
	}
	fmt.Println(sum)
}

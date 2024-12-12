package main

import (
	"AOC/h"
	"fmt"
	"maps"
	"slices"
	"strconv"
)

func main() {
	fmt.Println("Start")
	part1()
	//part2()
	part2diffrent()
}

func part1() {
	lines := h.GetLinesAsSlice()
	l := lines[0]
	n := h.ExtrapolateNumbersFromStringIgnoreNonDig(l)
	for i := 0; i < 25; i++ {
		newNums := make([]int, 0)
		for _, num := range n {
			if num == 0 {
				newNums = append(newNums, 1)
			} else if len(strconv.Itoa(num))%2 == 0 {
				str := []byte(strconv.Itoa(num))
				strs := []string{string(str[:len(str)/2]), string(str[len(str)/2:])}
				a, _ := strconv.Atoi(strs[0])
				b, _ := strconv.Atoi(strs[1])
				newNums = append(newNums, a)
				newNums = append(newNums, b)
			} else {
				newNums = append(newNums, num*2024)
			}
		}
		n = newNums
	}
	fmt.Println("part1, ", len(n))
}

func p2rec(element int, dep int, maxDep int) int {
	if dep == maxDep {
		return 1
	}
	newNums := make([]int, 0)
	if element == 0 {
		newNums = append(newNums, 1)
	} else if len(strconv.Itoa(element))%2 == 0 {
		str := []byte(strconv.Itoa(element))
		strs := []string{string(str[:len(str)/2]), string(str[len(str)/2:])}
		a, _ := strconv.Atoi(strs[0])
		b, _ := strconv.Atoi(strs[1])
		newNums = append(newNums, a)
		newNums = append(newNums, b)
	} else {
		newNums = append(newNums, element*2024)
	}
	sum := 0
	for _, e := range newNums {
		sum += p2rec(e, dep+1, maxDep)
	}
	return sum
}

func part2() {
	sum := 0
	lines := h.GetLinesAsSlice()
	l := lines[0]
	n := h.ExtrapolateNumbersFromStringIgnoreNonDig(l)

	for _, e := range n {
		sum += p2rec(e, 0, 75)
	}
	fmt.Println("part2, ", sum)
}

func part2diffrent() {
	lines := h.GetLinesAsSlice()
	sum := 0
	l := lines[0]
	amountOfElements := map[int]int{}
	n := h.ExtrapolateNumbersFromStringIgnoreNonDig(l)

	for _, nr := range n {
		amountOfElements[nr]++
	}

	for i := 0; i < 75; i++ {
		newAmounts := maps.Clone(amountOfElements)
		for number, amount := range amountOfElements {
			res := iteration(number)
			newAmounts[number] = newAmounts[number] - amount
			for _, r := range res {
				newAmounts[r] = newAmounts[r] + amount
			}
		}
		amountOfElements = newAmounts
	}
	for _, amount := range amountOfElements {
		sum += amount
	}
	fmt.Println("part1, ", sum)
}

var cache = map[int][]int{}

func iteration(e int) []int {
	r, ok := cache[e]
	if ok {
		return r
	}

	newNums := make([]int, 0)
	if e == 0 {
		newNums = append(newNums, 1)
	} else if len(strconv.Itoa(e))%2 == 0 {
		str := []byte(strconv.Itoa(e))
		strs := []string{string(str[:len(str)/2]), string(str[len(str)/2:])}
		a, _ := strconv.Atoi(strs[0])
		b, _ := strconv.Atoi(strs[1])
		newNums = append(newNums, a)
		newNums = append(newNums, b)
	} else {
		newNums = append(newNums, e*2024)
	}
	cache[e] = slices.Clone(newNums)
	return newNums
}

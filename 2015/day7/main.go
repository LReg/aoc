package main

import (
	"AOC/h"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Start")
	part1and2()
}

type O struct {
	dep              []string
	OptionalOperator string
}

var cache = map[string]int{}

func search(wire string, ops map[string]O, depth int) int {

	_, ok := cache[wire]
	if ok {
		return cache[wire]
	}

	if depth > len(ops)+5 {
		panic("loop?")
	}
	o := ops[wire]
	if o.OptionalOperator == "N" && unicode.IsDigit(rune(o.dep[0][0])) {
		n, _ := strconv.Atoi(o.dep[0])
		return n
	} else if o.OptionalOperator == "N" && !unicode.IsDigit(rune(o.dep[0][0])) {
		return search(o.dep[0], ops, depth+1)
	} else if o.OptionalOperator == "NOT" && !unicode.IsDigit(rune(o.dep[0][0])) {
		n := search(o.dep[0], ops, depth+1)
		cache[o.dep[0]] = n
		return int(^uint16(n))
	} else if o.OptionalOperator == "AND" {
		var n int
		var n1 int
		if unicode.IsDigit(rune(o.dep[0][0])) {
			n, _ = strconv.Atoi(o.dep[0])
		} else {
			n = search(o.dep[0], ops, depth+1)
			cache[o.dep[0]] = n
		}

		if unicode.IsDigit(rune(o.dep[1][0])) {
			n1, _ = strconv.Atoi(o.dep[1])
		} else {
			n1 = search(o.dep[1], ops, depth+1)
			cache[o.dep[0]] = n
		}

		return int(uint16(n) & uint16(n1))
	} else if o.OptionalOperator == "OR" {
		var n int
		var n1 int
		if unicode.IsDigit(rune(o.dep[0][0])) {
			n, _ = strconv.Atoi(o.dep[0])
		} else {
			n = search(o.dep[0], ops, depth+1)
			cache[o.dep[0]] = n
		}

		if unicode.IsDigit(rune(o.dep[1][0])) {
			n1, _ = strconv.Atoi(o.dep[1])
		} else {
			n1 = search(o.dep[1], ops, depth+1)
			cache[o.dep[0]] = n
		}
		return int(uint16(n) | uint16(n1))
	} else if o.OptionalOperator == "LSHIFT" {
		n := search(o.dep[0], ops, depth+1)
		n1, _ := strconv.Atoi(o.dep[1])
		return int(uint16(n) << n1)
	} else if o.OptionalOperator == "RSHIFT" {
		n := search(o.dep[0], ops, depth+1)
		cache[o.dep[0]] = n
		n1, _ := strconv.Atoi(o.dep[1])
		return int(uint16(n) >> n1)
	}
	fmt.Println("missed operator combination ?", wire, ops[wire])
	return -1
}

func part1and2() {
	lines := h.GetLinesAsSlice()
	ops := map[string]O{}
	wires := []string{}
	for _, l := range lines {
		p := strings.Split(l, "->")
		to := strings.TrimSpace(p[1])
		wires = append(wires, to)
		op := strings.TrimSpace(p[0])
		if strings.Contains(op, " ") {
			if strings.Contains(op, "NOT") {
				content := strings.Split(op, " ")
				o := O{[]string{content[1]}, content[0]}
				ops[to] = o
			} else {
				content := strings.Split(op, " ")
				o := O{[]string{content[0], content[2]}, content[1]}
				ops[to] = o
			}
		} else {
			ops[to] = O{[]string{op}, "N"}
		}
	}
	fmt.Println(ops)
	r := search("a", ops, 0)
	fmt.Println("Part 1 wire a: ", r)
	ops["b"] = O{
		dep:              []string{strconv.Itoa(r)},
		OptionalOperator: "N",
	}
	clear(cache)
	fmt.Println("Part 2 wire a: ", search("a", ops, 0))

}

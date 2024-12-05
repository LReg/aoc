package main

import (
	"AOC/h"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Rule struct {
	a int
	b int
}

func main() {
	fmt.Println("Start")
	//part1()
	part2()
}

func findRule(rules []Rule, f int) []Rule {
	res := make([]Rule, 0)
	for _, r := range rules {
		if r.a == f || r.b == f {
			res = append(res, r)
		}
	}
	return res
}

func part1() {
	sum := 0
	lines := h.GetLinesAsSlice()
	rules := make([]Rule, 0)
	pages := make([][]int, 0)
	goodPages := make([][]int, 0)

	stPages := 0
	for c, l := range lines {
		i := strings.Index(l, "|")
		if i == -1 {
			stPages = c
			break
		}
		one, two := h.StrSplitTwo(l, "|")
		c1, _ := strconv.Atoi(one)
		c2, _ := strconv.Atoi(two)
		rules = append(rules, Rule{c1, c2})
	}
	stPages++
	for i := stPages; i < len(lines); i++ {
		countsInt := h.ExtrapolateNumbersFromStringIgnoreNonDig(lines[i])
		pages = append(pages, countsInt)
	}

	//pages
	for _, p := range pages {
		good := true
		for pi, pnum := range p {
			rs := findRule(rules, pnum)
			for _, r := range rs {
				// page after
				for _, pnafter := range p[pi+1:] {
					if pnafter == r.a {
						good = false
						break
					}
				}
				for _, pbefore := range p[:pi] {
					if pbefore == r.b {
						good = false
						break
					}
				}

				if !good {
					break
				}
			}
		}
		if good {
			goodPages = append(goodPages, p)
		}
	}

	for _, p := range goodPages {
		median := p[(len(p) / 2)]
		fmt.Println(median)
		sum += median
	}

	fmt.Println(sum)

}

func isGood(p []int, rules []Rule) bool {
	good := true
	for pi, pnum := range p {
		rs := findRule(rules, pnum)
		for _, r := range rs {
			// page after
			for _, pnafter := range p[pi+1:] {
				if pnafter == r.a {
					good = false
					break
				}
			}
			for _, pbefore := range p[:pi] {
				if pbefore == r.b {
					good = false
					break
				}
			}

			if !good {
				break
			}
		}
	}
	return good
}

func part2() {
	sum := 0
	lines := h.GetLinesAsSlice()
	rules := make([]Rule, 0)
	pages := make([][]int, 0)
	badPages := make([][]int, 0)

	stPages := 0
	for c, l := range lines {
		i := strings.Index(l, "|")
		if i == -1 {
			stPages = c
			break
		}
		one, two := h.StrSplitTwo(l, "|")
		c1, _ := strconv.Atoi(one)
		c2, _ := strconv.Atoi(two)
		rules = append(rules, Rule{c1, c2})
	}
	stPages++
	for i := stPages; i < len(lines); i++ {
		countsInt := h.ExtrapolateNumbersFromStringIgnoreNonDig(lines[i])
		pages = append(pages, countsInt)
	}

	//pages
	for _, p := range pages {
		if !isGood(p, rules) {
			badPages = append(badPages, p)
		}
	}

	// sort pages
	for _, bp := range badPages {
		slices.SortFunc(bp, func(a int, b int) int {
			rs := findRule(rules, a)
			rs = append(rs, findRule(rules, b)...)
			for _, r := range rs {
				if a == r.a && b == r.b {
					return -1
				}
				if a == r.b && b == r.b {
					return 1
				}
			}
			return 0
		})
	}

	for _, p := range badPages {
		median := p[(len(p) / 2)]
		sum += median
	}

	fmt.Println(sum)
}

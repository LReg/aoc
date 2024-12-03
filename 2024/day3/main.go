package main

import (
	"AOC/h"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	//part1()
	//part2()
	alt()
}

func alt() {
	line := h.GetLinesAsOneString()
	for i, _ := range line {
		if strings.HasPrefix(line[i:], "mul(") {
			s := i + 4
			j := strings.Index(line[s:], ")")
			if j == -1 {
				continue
			}

			pts := strings.Split(line[s:j+s], ",")
			if len(pts) != 2 {
				continue
			}

			nums, failed := h.StrSlToIntSlSoftFail(pts)

			if failed {
				continue
			}

			fmt.Println(nums)
			// and so on

		}
	}
}

func part1() {
	fmt.Println("Start")
	sum := 0
	lines := h.GetLinesAsSlice()
	for _, l := range lines {
		res, err := regexp.Compile(`mul\((\d+),(\d+)\)`)
		if err != nil {
			continue
		}
		s := res.FindAllStringSubmatch(l, -1)
		for _, i := range s {
			a, _ := strconv.Atoi(i[1])
			b, _ := strconv.Atoi(i[2])
			erg := a * b

			sum += erg
		}
	}
	fmt.Println(sum)
}

func part2() {
	fmt.Println("Start")
	sum := 0
	lines := h.GetLinesAsSlice()
	l := ""
	for _, nl := range lines {
		l += nl
	}
	dos := strings.Split(l, "do()")

	for _, do := range dos {

		donsplit := strings.Split(do, "don't()")
		str := donsplit[0]

		res, _ := regexp.Compile(`mul\((\d+),(\d+)\)`)
		s := res.FindAllStringSubmatch(str, -1)
		for _, i := range s {
			a, _ := strconv.Atoi(i[1])
			b, _ := strconv.Atoi(i[2])
			erg := a * b

			sum += erg
		}
	}

	fmt.Println(sum)
}

package main

import (
	"AOC/helper"
	"fmt"
	"unicode"
)

func main() {
	part1()
	part2()
}

func part1() {
	fmt.Println("Start")
	sum := 0
	lines := helper.GetLinesAsRuneSlices()
	for _, sl := range lines {
		half := len(sl) / 2
		res := ' '
		for i := 0; i < half; i++ {
			for j := half; j < len(sl); j++ {
				if sl[i] == sl[j] {
					res = sl[i]
					break
				}
			}
			if res != ' ' {
				break
			}
		}
		p := prio(res)
		sum += p
	}

	fmt.Println(sum)
}

func prio(r rune) int {
	if unicode.IsUpper(r) {
		return int(r - 'A' + 27)
	} else {
		return int(r - 'a' + 1)
	}

}

func part2() {
	fmt.Println("Start")
	sum := 0
	lines := helper.GetLinesAsRuneSlices()

	for i := 0; i < len(lines)/3; i++ {
		item := ' '
		index := i * 3
		for _, a := range lines[index] {
			for _, b := range lines[index+1] {
				for _, c := range lines[index+2] {
					if a == b && b == c {
						item = a
						break
					}
				}
				if item != ' ' {
					break
				}
			}
			if item != ' ' {
				break
			}
		}
		sum += prio(item)
	}

	fmt.Println(sum)
}

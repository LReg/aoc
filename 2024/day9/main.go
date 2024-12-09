package main

import (
	"AOC/h"
	"fmt"
	"slices"
	"strconv"
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
		orig := []int{}
		used := true
		id := 0
		for _, c := range l {
			num, _ := strconv.Atoi(string(c))
			for i := 0; i < num; i++ {
				if used {
					orig = append(orig, id)

				} else {
					orig = append(orig, -1)
				}
			}

			used = !used
			if used {
				id++
			}
		}

		for i := len(orig) - 1; i >= 0; i-- {
			fillPos := slices.IndexFunc(orig, func(i int) bool { return i == -1 })
			if fillPos >= i {
				break
			}
			if orig[i] != -1 {
				orig[fillPos] = orig[i]
				orig[i] = -1
			}
		}

		for i, id := range orig {
			if id == -1 {
				break
			}
			sum += i * id
		}

	}
	fmt.Println(sum)
}

func part2() {
	sum := 0
	lines := h.GetLinesAsSlice()
	for _, l := range lines {
		orig := []int{}
		used := true
		id := 0
		for _, c := range l {
			num, _ := strconv.Atoi(string(c))
			for i := 0; i < num; i++ {
				if used {
					orig = append(orig, id)

				} else {
					orig = append(orig, -1)
				}
			}

			used = !used
			if used {
				id++
			}
		}

		for i := len(orig) - 1; i >= 0; i-- {
			numToMove := orig[i]
			if numToMove == -1 {
				continue
			}
			from := 0
			for j := i; j >= 0; j-- {
				if orig[j] != numToMove {
					from = j + 1
					break
				}
			}
			foundFree := 0
			startI := 0
			for put := 0; put < from; put++ {
				if orig[put] == -1 {
					foundFree++

					if foundFree == i-from+1 {
						countMoved := 0
						for k := startI; countMoved < i-from+1; k++ {
							orig[k] = numToMove
							countMoved++
						}

						for k := from; k <= i; k++ {
							orig[k] = -1
						}

						foundFree = 0
						break
					}

				} else {
					startI = put + 1
					foundFree = 0
				}
			}
			i -= i - from

		}

		for i, id := range orig {
			if id == -1 {
				continue
			}
			sum += i * id
		}

	}
	fmt.Println(sum)
}

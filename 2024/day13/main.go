package main

import (
	"AOC/h"
	"fmt"
	"os"
)

const (
	ACOST = 3
	BCOST = 1
)

func main() {
	fmt.Println("Start")
	//part1()
	part1dif()
	//part2()
	//part2diff()
	part2difdif()
}

func goalOrTooFar(pos h.Point, loc h.Point) int {
	if pos == loc {
		return 1
	}

	if pos.X > loc.X || pos.Y > loc.Y {
		return -1
	}

	return 0
}
func timesA(ta int, m Machine, part1 bool) (int, int) {
	cost := ta * ACOST
	p := h.Point{0, 0}
	p.X = m.AVector.X * ta
	p.Y = m.AVector.Y * ta
	r := goalOrTooFar(p, m.PriceLocation)
	if r != 0 {
		return r, cost
	}
	bPresses := (m.PriceLocation.X - p.X) / m.BVector.X
	wouldFitY := (m.PriceLocation.Y-p.Y)%m.BVector.Y == 0
	wouldFitX := (m.PriceLocation.X-p.X)%m.BVector.X == 0
	p.X += m.BVector.X * bPresses
	p.Y += m.BVector.Y * bPresses
	if !wouldFitX || !wouldFitY || (part1 && bPresses > 100) || goalOrTooFar(p, m.PriceLocation) != 1 {
		return -1, cost
	} else {
		return 1, cost + bPresses*BCOST
	}
}

func part1dif() {
	sum := 0
	lines := h.GetLinesAsSlice()
	machines := []Machine{}
	for i := 0; i < len(lines); {
		m := Machine{}
		for j := 0; j < 3; j++ {
			if j == 0 {
				nms := h.ExtrapolateNumbersFromStringIgnoreNonDig(lines[j+i])
				m.AVector.X = nms[0]
				m.AVector.Y = nms[1]
			} else if j == 1 {
				nms := h.ExtrapolateNumbersFromStringIgnoreNonDig(lines[j+i])
				m.BVector.X = nms[0]
				m.BVector.Y = nms[1]
			} else {
				nms := h.ExtrapolateNumbersFromStringIgnoreNonDig(lines[j+i])
				m.PriceLocation.X = nms[0]
				m.PriceLocation.Y = nms[1]
			}
		}
		machines = append(machines, m)
		i += 4
	}
	for _, m := range machines {
		for tA := 0; ; tA++ {
			status, cost := timesA(tA, m, true)
			if status == 1 {
				sum += cost
				break
			}
			if tA > 100 || (tA*m.AVector.X > m.PriceLocation.X && tA*m.AVector.Y > m.PriceLocation.Y) {
				break
			}
		}
	}
	fmt.Println("part 1", sum)
}

type Machine struct {
	AVector, BVector h.Point
	PriceLocation    h.Point
}

func createnm(m *Machine, currentPoint h.Point, cost int) int {
	if currentPoint == m.PriceLocation {
		return cost
	}

	if currentPoint.X > m.PriceLocation.X || currentPoint.Y > m.PriceLocation.Y {
		return -1
	}

	pointAfterVB := h.Point{currentPoint.X + m.BVector.X, currentPoint.Y + m.BVector.Y}

	wb := createnm(m, pointAfterVB, cost+BCOST)
	if wb != -1 {
		return wb
	}

	pointAfterVA := h.Point{currentPoint.X + m.AVector.X, currentPoint.Y + m.AVector.Y}
	wa := createnm(m, pointAfterVA, cost+ACOST)
	if wa != -1 {
		return wa
	}

	return -1
}

func part1() {
	sum := 0
	lines := h.GetLinesAsSlice()
	machines := []Machine{}
	for i := 0; i < len(lines); {
		m := Machine{}
		for j := i; j < i+3; j++ {
			if j == 0 {
				nms := h.ExtrapolateNumbersFromStringIgnoreNonDig(lines[j])
				m.AVector.X = nms[0]
				m.AVector.Y = nms[1]
			} else if j == 1 {
				nms := h.ExtrapolateNumbersFromStringIgnoreNonDig(lines[j])
				m.BVector.X = nms[0]
				m.BVector.Y = nms[1]
			} else {
				nms := h.ExtrapolateNumbersFromStringIgnoreNonDig(lines[j])
				m.PriceLocation.X = nms[0]
				m.PriceLocation.Y = nms[1]
			}
		}
		machines = append(machines, m)
		i += 4
	}
	for _, m := range machines {
		weight := createnm(&m, h.Point{0, 0}, 0)
		sum += weight
	}
	fmt.Println(sum)
}

func part2() {
	sum := float64(0)
	lines := h.GetLinesAsSlice()
	machines := []Machine{}
	for i := 0; i < len(lines); {
		m := Machine{}
		for j := 0; j < 3; j++ {
			if j == 0 {
				nms := h.ExtrapolateNumbersFromStringIgnoreNonDig(lines[j+i])
				m.AVector.X = nms[0]
				m.AVector.Y = nms[1]
			} else if j == 1 {
				nms := h.ExtrapolateNumbersFromStringIgnoreNonDig(lines[j+i])
				m.BVector.X = nms[0]
				m.BVector.Y = nms[1]
			} else {
				nms := h.ExtrapolateNumbersFromStringIgnoreNonDig(lines[j+i])
				m.PriceLocation.X = nms[0] * 10000000000000
				m.PriceLocation.Y = nms[1] * 10000000000000
			}
		}
		machines = append(machines, m)
		i += 4
	}
	for _, m := range machines {
		fmt.Println("Machine: a, b, price", m.AVector, m.BVector, m.PriceLocation)
		ma := float64(m.AVector.Y) / float64(m.AVector.X)
		mb := float64(m.BVector.Y) / float64(m.BVector.X)
		fmt.Println("ma, mb", ma, mb)
		na := float64(0)
		nb := ((mb * float64(m.PriceLocation.X)) + float64(m.PriceLocation.Y)) * -1
		fmt.Println("na, nb", na, nb)
		// Berechnung des Schnittpunkts
		if ma != mb {
			x := (nb - na) / (ma - mb)
			y := ma*x + na
			fmt.Println("x, y", x, y)
			ta := x / float64(m.AVector.X)
			tb := float64(m.PriceLocation.X) - x/float64(m.BVector.X)
			fmt.Println("ta, tb", ta, tb)
			sum += ta*float64(ACOST) + tb*float64(BCOST)
		} else {
			fmt.Println("parralel")
			// paralel -> works only with b
			tb := float64(m.PriceLocation.X) / float64(m.BVector.X)
			sum += tb * float64(BCOST)
		}
	}
	fmt.Println("part 2", sum)
}

func part2diff() {
	sum := float64(0)
	lines := h.GetLinesAsSlice()
	machines := []Machine{}
	for i := 0; i < len(lines); {
		m := Machine{}
		for j := 0; j < 3; j++ {
			if j == 0 {
				nms := h.ExtrapolateNumbersFromStringIgnoreNonDig(lines[j+i])
				m.AVector.X = nms[0]
				m.AVector.Y = nms[1]
			} else if j == 1 {
				nms := h.ExtrapolateNumbersFromStringIgnoreNonDig(lines[j+i])
				m.BVector.X = nms[0]
				m.BVector.Y = nms[1]
			} else {
				nms := h.ExtrapolateNumbersFromStringIgnoreNonDig(lines[j+i])
				m.PriceLocation.X = nms[0] * 10000000000000
				m.PriceLocation.Y = nms[1] * 10000000000000
			}
		}
		machines = append(machines, m)
		i += 4
	}
	for _, m := range machines {
		x, _, ok := intersection(float64(m.PriceLocation.X), float64(m.PriceLocation.Y), float64(m.AVector.X), float64(m.AVector.Y), float64(m.BVector.X), float64(m.BVector.Y))
		if !ok {
			fmt.Println("not ok")
		}
		ta := x / float64(m.AVector.X)
		tb := float64(m.PriceLocation.X) - x/float64(m.BVector.X)
		fmt.Println("ta, tb", ta, tb)
		sum += ta*float64(ACOST) + tb*float64(BCOST)
	}
	fmt.Println("part 2", sum)
}
func intersection(
	x, y, dx1, dy1, dx2, dy2 float64,
) (float64, float64, bool) {
	denom := dx1*dy2 - dy1*dx2
	if denom == 0 {
		return 0, 0, false
	}

	t := (x*dy2 - y*dx2) / denom

	px := t * dx1
	py := t * dy1

	fmt.Println(px, py)
	return px, py, true
}

func part2difdif() {
	sum := 0
	lines := h.GetLinesAsSlice()
	machines := []Machine{}
	for i := 0; i < len(lines); {
		m := Machine{}
		for j := 0; j < 3; j++ {
			if j == 0 {
				nms := h.ExtrapolateNumbersFromStringIgnoreNonDig(lines[j+i])
				m.AVector.X = nms[0]
				m.AVector.Y = nms[1]
			} else if j == 1 {
				nms := h.ExtrapolateNumbersFromStringIgnoreNonDig(lines[j+i])
				m.BVector.X = nms[0]
				m.BVector.Y = nms[1]
			} else {
				nms := h.ExtrapolateNumbersFromStringIgnoreNonDig(lines[j+i])
				m.PriceLocation.X = nms[0] + 10000000000000
				m.PriceLocation.Y = nms[1] + 10000000000000
			}
		}
		machines = append(machines, m)
		i += 4
	}
	for _, m := range machines {
		for tA := 0; ; tA++ {
			p := h.Point{m.AVector.X * tA, m.AVector.Y * tA}
			if (m.PriceLocation.X-p.X)%m.BVector.X == 0 && (m.PriceLocation.Y-p.Y)%m.BVector.Y == 0 {
				if (m.PriceLocation.X-p.X)/m.BVector.X == (m.PriceLocation.Y-p.Y)/m.BVector.Y {
					fmt.Println("Found")
					os.Exit(1)

				} else {
					d := h.Abs((m.PriceLocation.X-p.X)/m.BVector.X - (m.PriceLocation.Y-p.Y)/m.BVector.Y)
					fmt.Println("diff", h.Abs((m.PriceLocation.X-p.X)/m.BVector.X-(m.PriceLocation.Y-p.Y)/m.BVector.Y))
					if d > 220417217 {
						tA += d / 10000
					}
				}

				//fmt.Println("fit machine ", i, "times a", tA)
				//bPresses := (m.PriceLocation.X - p.X) / m.BVector.X
				//p.X += m.BVector.X * bPresses
				//p.Y += m.BVector.Y * bPresses
			}
		}
		fmt.Println("part 2", sum)
	}
}

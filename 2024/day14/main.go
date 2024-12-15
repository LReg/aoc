package main

import (
	"AOC/h"
	"fmt"
	"slices"
	"strings"
)

func main() {
	fmt.Println("Start")
	//part1()
	part2()
}

type Robot struct {
	pos h.Point
	vel h.Point
}

const (
	H       = 103
	W       = 101
	SECONDS = 1
)

func part1() {
	sum := 1
	lines := h.GetLinesAsSlice()
	robots := []Robot{}

	for _, l := range lines {
		spl := strings.Split(l, " ")
		p := h.ExtrapolateNumbersFromStringIgnoreNonDig(spl[0])
		v := h.ExtrapolateNumbersFromString(strings.Split(spl[1], "=")[1], ",")
		robots = append(robots, Robot{
			h.Point{
				p[0],
				p[1],
			},
			h.Point{
				v[0],
				v[1],
			},
		})
	}

	for i, r := range robots {
		x := r.pos.X + r.vel.X*SECONDS
		y := r.pos.Y + r.vel.Y*SECONDS
		if x < 0 {
			x = (W - (h.Abs(x) % W)) % W
		} else {
			x = x % W
		}
		if y < 0 {
			y = (H - (h.Abs(y) % H)) % H
		} else {
			y = y % H
		}
		r.pos = h.Point{
			X: x,
			Y: y,
		}
		robots[i] = r
	}

	g := h.CreateGrid(W, H)
	g.ForEachPoint(func(p h.Point) {
		g.Set(p, '.')
	})
	for _, r := range robots {
		g.Set(r.pos, 'R')
	}
	h.PrintGrid(g)

	XMID := (W / 2) + 1
	YMID := (H / 2) + 1

	for QX := 0; QX < 2; QX++ {
		for QY := 0; QY < 2; QY++ {
			counter := 0
			for _, r := range robots {
				if r.pos.X > QX*(XMID)-1 && r.pos.X < (QX*XMID)+XMID-1 &&
					r.pos.Y > QY*(YMID)-1 && r.pos.Y < (QY*YMID)+YMID-1 {
					counter++
				}
			}
			sum *= counter
		}
	}

	fmt.Println(robots)
	fmt.Println(sum)
}

func part2() {
	lines := h.GetLinesAsSlice()
	robots := []Robot{}

	for _, l := range lines {
		spl := strings.Split(l, " ")
		p := h.ExtrapolateNumbersFromStringIgnoreNonDig(spl[0])
		v := h.ExtrapolateNumbersFromString(strings.Split(spl[1], "=")[1], ",")
		robots = append(robots, Robot{
			h.Point{
				p[0],
				p[1],
			},
			h.Point{
				v[0],
				v[1],
			},
		})
	}

	for round := 0; round < 15000; round++ {
		fmt.Println(round)
		for i, r := range robots {
			x := r.pos.X + r.vel.X*SECONDS
			y := r.pos.Y + r.vel.Y*SECONDS
			if x < 0 {
				x = (W - (h.Abs(x) % W)) % W
			} else {
				x = x % W
			}
			if y < 0 {
				y = (H - (h.Abs(y) % H)) % H
			} else {
				y = y % H
			}
			r.pos = h.Point{
				X: x,
				Y: y,
			}
			robots[i] = r
		}
		points := h.Map(robots, func(r Robot) h.Point { return r.pos })
		foundSuspected := false
		for _, p := range points {
			neis := p.Neighbours()
			allFound := true
			for _, n := range neis {
				if !slices.Contains(points, n) {
					allFound = false
					break
				}
			}
			if allFound {
				fmt.Println(p)
				foundSuspected = true
				break
			}
		}
		if foundSuspected {
			break
		}
	}

	g := h.CreateGrid(W, H)
	g.ForEachPoint(func(p h.Point) {
		g.Set(p, '.')
	})
	for _, r := range robots {
		g.Set(r.pos, 'R')
	}
	h.PrintGrid(g)
}

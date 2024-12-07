package main

import (
	"AOC/h"
	"fmt"
	"slices"
)

func main() {
	extrapolateNumbers()
	compareRunes()
	compareRunesComplex()
	streams()
	stack()
	TestDijkstra()
	TestDijkstra2()
	println("Tests run successful")
}

func stack() {
	stack := make([]int, 0)
	h.Push(&stack, 1)
	h.Push(&stack, 2)
	_ = h.Pop(&stack)
	_ = stack[len(stack)-1]
	stack = stack[:len(stack)-1]
}

func streams() {
	strs := []string{"123", " 456", "789"}
	n, _ := h.StrSlToIntSlSoftFail(strs)
	if n[0] != 123 || n[1] != 456 || n[2] != 789 {
		panic("fail")
	}
	n = h.Map(strs, h.MapStrToInt)
	if n[0] != 123 || n[1] != 456 || n[2] != 789 {
		panic("fail")
	}
	sum := h.SumIntSl(n)
	if sum != 1368 {
		panic("fail")
	}
	rev := h.Map(strs, h.MapReverseString)
	if rev[0] != "321" || rev[1] != "654 " || rev[2] != "987" {
		panic("fail")
	}
}

func compareRunesComplex() {
	t1 := []rune("12312312313HERE")
	found := false
	for i, _ := range t1 {
		/**
		 this fails
		if string(t1[i:]) == "HERE" {
			println("here")
		}
		*/

		if h.SafeCompRuneSl(t1, i, "HERE") {
			found = true
		}
	}
	if !found {
		panic("fail")
	}

	t2 := []rune("12312312313HERE")
	found = false
	for i, _ := range t2 {
		/**
		 this fails
		if string(t1[i:]) == "HERE" {
			println("here")
		}
		*/

		if h.SafeCompRuneSl(t2, i, "HERE") {
			found = true
		}
	}
	if !found {
		panic("fail")
	}
}

func compareRunes() {
	t1 := []rune("123")
	res := h.SafeCompRuneSl(t1, 0, "123")
	if !res {
		panic("fail")
	}
	res = h.SafeCompRuneSl(t1, 0, "1234")
	if res {
		panic("fail")
	}
	res = string(t1) == "123"
	if !res {
		panic("fail")
	}
	res = string(t1) == "1234"
	if res {
		panic("fail")
	}
}

func extrapolateNumbers() {
	t1 := "12 35 65"
	res := h.ExtrapolateNumbersFromString(t1, " ")
	if res[0] != 12 || res[1] != 35 || res[2] != 65 {
		panic("fail")
	}
	t2 := "12  35 65"
	res = h.ExtrapolateNumbersFromString(t2, " ")
	if res[0] != 12 || res[1] != 35 || res[2] != 65 {
		panic("fail")
	}
	t3 := "12,  35, 65"
	res = h.ExtrapolateNumbersFromString(t3, ",")
	if res[0] != 12 || res[1] != 35 || res[2] != 65 {
		panic("fail")
	}
	t4 := "24778   15223"
	res = h.ExtrapolateNumbersFromString(t4, " ")
	if res[0] != 24778 || res[1] != 15223 {
		panic("fail")
	}
	t5 := "(24778   15223)"
	res = h.ExtrapolateNumbersFromStringIgnore(t5, " ", []string{"(", ")"})
	if res[0] != 24778 || res[1] != 15223 {
		panic("fail")
	}
	t6 := "(24778   15223)"
	res = h.ExtrapolateNumbersFromStringIgnoreNonDig(t6)
	if res[0] != 24778 || res[1] != 15223 {
		panic("fail")
	}
}

func TestDijkstra() {
	lines := make([]string, 0)
	lines = append(lines, "111")
	lines = append(lines, "991")
	lines = append(lines, "991")

	grid := h.ConvertLinesToGrid(lines)

	// Testfälle vorbereiten
	start := h.Point{X: 0, Y: 0}
	end := h.Point{X: 2, Y: 2}

	// Erwartete Ergebnisse
	expectedPath := []h.Point{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 2, Y: 1}, {X: 2, Y: 2}}
	expectedDistance := 4

	// Funktion testen
	path, distance := grid.DijkstraPosNum(start, end)

	p := slices.EqualFunc(expectedPath, path, h.EqualsPoint)
	d := distance == expectedDistance
	if !p || !d {
		panic("Dijksta failed")
	}
}

func TestDijkstra2() {
	lines := make([]string, 0)
	lines = append(lines, "121")
	lines = append(lines, "191")
	lines = append(lines, "111")

	grid := h.ConvertLinesToGrid(lines)

	// Testfälle vorbereiten
	start := h.Point{X: 0, Y: 0} // Startpunkt: Unten links
	end := h.Point{X: 2, Y: 2}   // Endpunkt: Oben rechts

	// Erwartete Ergebnisse
	expectedPath := []h.Point{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}, {X: 1, Y: 2}, {X: 2, Y: 2}}
	expectedDistance := 4

	// Funktion testen
	path, distance := grid.DijkstraPosNum(start, end)
	fmt.Println(path, distance)

	p := slices.EqualFunc(expectedPath, path, h.EqualsPoint)
	d := distance == expectedDistance
	if !p || !d {
		panic("Dijkstra failed in second case")
	}
}

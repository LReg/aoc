package h

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetLinesAsSlice() []string {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	slice := make([]string, 0, 50)
	for scanner.Scan() {
		slice = append(slice, scanner.Text())
	}
	return slice
}

func GetLinesAsRuneSlices() [][]rune {
	lines := GetLinesAsSlice()
	lr := make([][]rune, 0)
	for _, l := range lines {
		ra := []rune(l)
		lr = append(lr, ra)
	}
	return lr
}

func GetLinesAsOneRuneSlice() []rune {
	lines := GetLinesAsSlice()
	arr := make([]rune, 0)
	for _, l := range lines {
		arr = append(arr, []rune(l)...)
	}
	return arr
}

func GetLinesAsOneString() string {
	lines := GetLinesAsSlice()
	return strings.Join(lines, "")
}

func PrintGrid(grid [][]byte) {
	// zeile
	for y, _ := range Iter(len(grid[0])) {
		//spalte
		for x, _ := range Iter(len(grid)) {
			fmt.Printf("%c", grid[x][y])
		}
		fmt.Printf("\n")
	}
}

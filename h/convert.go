package h

// ConvertLinesToGrid
// access grid[x][y] x -> horizontal, y -> vertical
// remember to first iterate over one
//
// for y, _ := range h.Iter(len(grid[0])) { ->zeile
//
//	for x, _ := range h.Iter(len(grid)) { ->spalte
func ConvertLinesToGrid(lines []string) Grid {
	longestLine := Reduce(lines, func(prev int, curr string) int {
		l := len(curr)
		if l > prev {
			return l
		} else {
			return prev
		}
	}, 0)

	grid := make([][]byte, longestLine)
	for i, _ := range grid {
		grid[i] = make([]byte, len(lines))
	}

	for i, _ := range lines {
		for j, _ := range lines[i] {
			grid[j][i] = lines[i][j]
		}
	}

	return grid
}

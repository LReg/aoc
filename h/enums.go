package h

const (
	EAST = iota
	NORTH
	NORTHEAST
	SOUTHEAST
	WEST
	SOUTH
	NORTHWEST
	SOUTHWEST
)

func GetAllDirs() []int {
	return []int{EAST, NORTH, NORTHEAST, SOUTHEAST, WEST, SOUTH, NORTHWEST, SOUTHWEST}
}

func GetBasicDirs() []int {
	return []int{EAST, NORTH, WEST, SOUTH}
}

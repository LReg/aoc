package main

import (
	"AOC/helper"
)

func main() {
	extrapolateNumbers()
	println("Tests run successful")
}

func extrapolateNumbers() {
	t1 := "12 35 65"
	res := helper.ExtrapolateNumbersFromString(t1, " ")
	if res[0] != 12 || res[1] != 35 || res[2] != 65 {
		panic("fail")
	}
	t2 := "12  35 65"
	res = helper.ExtrapolateNumbersFromString(t2, " ")
	if res[0] != 12 || res[1] != 35 || res[2] != 65 {
		panic("fail")
	}
	t3 := "12,  35, 65"
	res = helper.ExtrapolateNumbersFromString(t3, ",")
	if res[0] != 12 || res[1] != 35 || res[2] != 65 {
		panic("fail")
	}
	t4 := "24778   15223"
	res = helper.ExtrapolateNumbersFromString(t4, " ")
	if res[0] != 24778 || res[1] != 15223 {
		panic("fail")
	}
}

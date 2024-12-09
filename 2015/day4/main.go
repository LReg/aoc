package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Start")
	part1()
	part2()
}

func part1() {
	sc := "bgvyzdsv"
	for i := 1; true; i++ {
		str := strconv.Itoa(i)
		res := md5.Sum([]byte(sc + str))
		s := hex.EncodeToString(res[:])
		allGood := true
		for j := 0; j < 6; j++ {
			if '0' != s[j] {
				allGood = false
			}
		}
		if allGood {
			fmt.Println("res:", i)
			break
		}
	}
}

func part2() {
	sc := "bgvyzdsv"
	for i := 1; true; i++ {
		str := strconv.Itoa(i)
		res := md5.Sum([]byte(sc + str))
		s := hex.EncodeToString(res[:])
		allGood := true
		for j := 0; j < 5; j++ {
			if '0' != s[j] {
				allGood = false
			}
		}
		if allGood {
			fmt.Println("res:", i)
			break
		}
	}
}

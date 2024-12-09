package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	res := md5.Sum([]byte("abcdef609043"))
	fmt.Println(str)
}

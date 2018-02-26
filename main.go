package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world!")
	fmt.Println(calc(2, 3))
}

func calc(i, j int) (k int) {
	k = i + j
	return
}

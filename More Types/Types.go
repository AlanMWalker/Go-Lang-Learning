package main

import (
	"fmt"
)

func main() {
	var num int
	var nump *int
	num = 10
	nump = &num

	fmt.Printf("%d %d \n", nump, num)
}

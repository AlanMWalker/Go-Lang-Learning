package main

import "fmt"

const MAX_NUMBERS = 10

func main() {
	var number float32
	sum := float32(0.0)

	fmt.Printf("Enter %d numbers\n", MAX_NUMBERS)
	i := 0
	for i < int(MAX_NUMBERS) {
		fmt.Scan(&number)
		sum += number
		i += 1
	}
	average := sum / float32(MAX_NUMBERS)
	fmt.Printf("Average = %f\n", average)
}

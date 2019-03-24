package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	num := rand.Int() % 100
	var guess int

	for {
		fmt.Printf("Guess a number between 1 - 100\n")
		fmt.Scanf("%d\n", &guess)

		switch {
		case guess > num:
			fmt.Println("Too high")
		case guess == num:
			fmt.Println("You win")
			return
		case guess < num:
			fmt.Println("Too low")
		}
	}
}

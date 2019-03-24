package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func basicFor() {
	//basic for
	for i := 0; i < 10; i++ {
		fmt.Println("Basic For ", i)
	}
}

func omittedFor() {
	i := 0
	for i < 10 {
		i += 1
		fmt.Println("Omitted For ", i)
	}
}

func foreverLoopy() {
	for {
	}
}

func printMe() {
	fmt.Print("I should be printed last")
}

func main() {
	go foreverLoopy()
	go basicFor()
	go omittedFor()

	// defer -> defers execution of a function until the surrounding function returns
	// if you have multiple defers they are placed on a stack and popped out last-in-first-out order
	defer printMe()
	// basic if
	x := float64(100.0)
	if x > 0 {
		fmt.Println("Basic if")
	} else {
		fmt.Println("else case")
	}

	// if with pre-execution statement
	if y := math.Max(x, 40); x > 0 && y > 0 {
		fmt.Println("Pre execution")
	}

	// switch
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// switch (true) (good for long if else chains)
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	// for concurrency
	time.Sleep(1000 * time.Millisecond)
}

package main

import (
	"fmt"
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

func main() {
	go foreverLoopy()
	go basicFor()
	go omittedFor()

	// for concurrency
	time.Sleep(5000 * time.Millisecond)
}

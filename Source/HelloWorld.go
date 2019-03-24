package main

import "fmt"

// normal functions
// func add(x int, y int ){
func add(x, y int) int {
	return x + y
}

func doesNothingVoid() {
	fmt.Println("Does nothing at all")
	return
}

// Multiple go func returns
func addSub(x, y int) (int, int) {
	return x + y, y - x
}

func addSubModified(x, y int) (add, sub int) {
	add = x + y
	sub = x - y
	return
}

/*
GO LANG BASIC TYPES LIST
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
*/

func main() {
	doesNothingVoid()
	fmt.Println("Normal function ", add(1, 0))
	fmt.Print("Multi return")
	fmt.Println(addSub(1, 2))
	fmt.Print("Multi return named ")
	fmt.Println(addSubModified(1, 2))
	var a string
	a = "Hello world"
	a += "h"
	fmt.Println(a)
	/*
		Variable decleration
	*/
	var i int
	x := float64(10.0)
	fmt.Println(i)
	fmt.Println(x)
	fmt.Scanf("%d", &i)
	fmt.Println("Value of i ", i)
}

package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

type StrPair struct {
	a string
	b string
}

type Vertex struct {
	x, y int32
}

type FVertex struct {
	x, y float32
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// function closure example
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	/* OS Args with range for loop*/
	for i := range os.Args {
		fmt.Printf("Args %d %s \n", (i + 1), os.Args[i])
	}
	rand.Seed(time.Now().UnixNano())
	var num int32
	var numPtr *int32 = nil // ptr in go lang
	num = 10

	if numPtr == nil {
		fmt.Printf("%d\n", numPtr)
	}

	numPtr = &num

	fmt.Printf("%d %d \n", numPtr, num)

	//	Struct init example
	v := Vertex{1, 1}
	fmt.Println("Vertex print ", Vertex{1, 1})
	fmt.Println("Vertex print ", v)
	fmt.Println("Vertex print ", StrPair{"hello", "world"})

	//	Struct member access example
	fmt.Printf("Vertex individual members %d %d \n", v.x, v.y)

	//	Struct pointers
	vertPtr := &v
	fmt.Printf("Vertex ptr members %d %d \n", vertPtr.x, vertPtr.y)

	//	Arrays
	var a [15]int32
	for i := range a {
		a[i] = int32(rand.Int()) % 100
		fmt.Printf("Array member print %d\n", a[i])
	}

	// Slices
	var slice []int32 = a[0:5]
	fmt.Println("Slice print ", slice)
	/*
		//For the array
		var a [10]int
		//these slice expressions are equivalent:
		a[0:10]
		a[:10]
		a[0:]
		a[:]
	*/
	boolSlice := []bool{true, false, false}
	fmt.Println(boolSlice)
	sliceSize := len(slice)
	sliceCapacity := cap(slice)

	fmt.Printf("Capacity %d - Length %d\n", sliceSize, sliceCapacity)

	// make()
	makeSliceA := make([]int32, 5) // len(makeSliceA) = 5
	fmt.Printf("MakeSliceA cap = %d len = %d \n", cap(makeSliceA), len(makeSliceA))
	fmt.Println(makeSliceA)

	makeSliceB := make([]int32, 5, 1000)
	fmt.Printf("MakeSliceB cap = %d len = %d \n", cap(makeSliceB), len(makeSliceB))
	fmt.Println(makeSliceB)
	// Slices of slices
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	fmt.Println(board)

	makeSliceB = append(makeSliceB, 0)
	fmt.Printf("MakeSliceB (post-append) cap = %d len = %d \n", cap(makeSliceB), len(makeSliceB))
	/*
		//You can skip the index or value by assigning to _.

		for i, _ := range pow
		for _, value := range pow

		//If you only want the index, you can omit the second variable.

		for i := range pow
	*/
	mapExample := make(map[string]Vertex)
	mapExample["TestMapEntry"] = Vertex{66, 6}
	fmt.Println(mapExample)
	//	Map literal
	var m = map[string]FVertex{
		"Bell Labs": FVertex{
			40.68433, -74.39967,
		},
		"Google": FVertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)
	//insert/update
	m["Google"] = FVertex{1.0, 1.0}
	fmt.Println(m)
	//access
	elem := m["Google"]
	fmt.Println(m, elem)
	//delete
	delete(m, "Google")
	fmt.Println(m)

	/* Function Values*/
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println("Func val ", compute(hypot))
	return
}

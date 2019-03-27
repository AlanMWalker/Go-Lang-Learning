package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float32
}

// by-value normal method (more like const class func in C++)
func (v Vertex) DotProduct(otherVertex Vertex) float32 {
	return ((v.x * otherVertex.x) + (v.y * otherVertex.y))
}

// method for a type

type f32 float32

func (f f32) abs() f32 {
	if f < 0 {
		return f * f32(-1.0)
	}
	return f
}

// pointer function which can affect an instances members (more like a non-const class func in C++)
func (v *Vertex) AbsValues() {
	v.x = float32(math.Abs(float64(v.x)))
	v.y = float32(math.Abs(float64(v.y)))
}

// interfaces
type Printer interface {
	print()
}

type str string

func (s str) print() {
	fmt.Print(s)
}

type Poly struct {
	vertices [100]Vertex
}

// by pointer to avoid copies
func (p *Poly) print() {

	// account for nil receiever (instead of null ptr exception)
	if p == nil {
		fmt.Print("Nill reciever passed to func\n")
		return
	}
	for i := range p.vertices {
		p.vertices[i] = Vertex{0, 1}
	}
	fmt.Println(p)
}

func main() {
	// methods
	v := Vertex{1.0, -1.0}
	fmt.Println("Dot product", v.DotProduct(Vertex{1, 2}))
	fmt.Println("V before abs ", v)
	v.AbsValues()
	fmt.Println("V after abs ", v)
	fmt.Print("f32 member ", f32(-100).abs())

	// interfaces
	var p Printer
	var s str
	var ply Poly
	p = s
	p.print()
	p = &ply // if not ptr won't work because Poly print method requires pointer for its print interface definition
	p.print()

	// TODO : Type assertions page 15 out of 26 on Go lang interfaces
}

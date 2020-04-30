package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float32
}

type square struct {
	side float32
}

type rec struct {
	length, width float32
}

type shape interface {
	area() float32
}

func (c circle) area() float32 {
	return 2 * math.Pi * c.radius * c.radius
}

func (s square) area() float32 {
	return s.side * s.side
}

func (r rec) area() float32 {
	return 2 * (r.length + r.width)
}

func calcArea(s shape) float32 {
	return s.area()
}

func main() {

	c := circle{
		4,
	}

	s := square{
		4,
	}

	r := rec{
		4, 6,
	}

	var sh shape = circle{
		9,
	}

	fmt.Printf("%T", sh)
	fmt.Println()
	fmt.Println(calcArea(c))
	fmt.Println(calcArea(s))
	fmt.Println(calcArea(r))

}

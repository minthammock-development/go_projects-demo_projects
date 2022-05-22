package main

import "fmt"
import "math"

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect ) area() float64 {
	return r.height * r.width
}

func (r rect) perim() float64 {
	return r.height*2 + r.width*2
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c circle) perim() float64 {
	return math.Pi * 2 * c.radius
}

func measure (g geometry) {
	fmt.Println(g)
	fmt.Println(g.perim())
	fmt.Println(g.area())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}


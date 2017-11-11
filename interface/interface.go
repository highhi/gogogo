package main

import(
	"math"
	"fmt"
)

type geometory interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (rect rect) area() float64 {
	return rect.width * rect.height
}

func (rect rect) perim() float64 {
	return 2 * rect.width + 2 * rect.height
}

func (circle circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func (circle circle) perim() float64 {
	return 2 * math.Pi * circle.radius
}

func measure(g geometory)  {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
package main

import "fmt"

type circle struct {
	radius float64
}

func (circle circle) area() float64 {
	return 3.14 * circle.radius * circle.radius
}
func main() {
	var c circle
	c.radius = 4.0
	fmt.Println("area: ", c.area())
}

package main

import (
	"fmt"
	"math"
)

//求ax^2 + bx + c = 0的解
//input a b c
//if b^2 - 4ac > 0  2个解
//if b^2 - 4ac = 0  1个解
//if b^2 - 4ac < 0  0个解
//x1 = (-b + sqrt(b2-4ac)) / 2a       x2 = (-b - sqrt(b2-4ac)) / 2a

func main() {
	var a, b, c float64
	a = 3.0
	b = 100.0
	c = 6.0
	//fmt.Scanln(&a, &b, &c)
	if math.Pow(b, 2)-4*a*c > 0 {
		fmt.Println("The answer x1 is:\t", (-b+math.Sqrt(b*b-4*a*c))/(2*a),
			"\nThe answer x2 is:\t", (-b-math.Sqrt(b*b-4*a*c))/(2*a))
	} else if math.Pow(b, 2)-4*a*c == 0 {
		fmt.Println("\nThe answer x1 is:\t", (-b+math.Sqrt(b*b-4*a*c))/(2*a))
	} else {
		fmt.Println("There is no answer!")
	}
}

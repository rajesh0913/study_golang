package main

import "fmt"

type Calculator struct {
	num1 float64
	num2 float64
}

func (cal *Calculator) Addition() float64 {
	return (*cal).num1 + (*cal).num2
}
func (cal *Calculator) Sub() float64 {
	return (*cal).num1 - (*cal).num2
}
func (cal *Calculator) Multi() float64 {
	return (*cal).num1 * (*cal).num2
}
func (cal *Calculator) Devid() float64 {
	return (*cal).num1 / (*cal).num2
}
func (cal *Calculator) gerRes(operator byte) float64 {
	res := 0.0
	switch operator {
	case '+':
		res = (*cal).num1 + (*cal).num2
	case '-':
		res = (*cal).num1 - (*cal).num2
	case '*':
		res = (*cal).num1 * (*cal).num2
	case '/':
		res = (*cal).num1 / (*cal).num2
	default:
		fmt.Println("输入错误！")
	}
	return res
}

func main() {
	var Cal Calculator
	Cal.num1 = 10.0
	Cal.num2 = 20.0
	fmt.Printf("Num1 + Num2 = %v\n", fmt.Sprintf("%.2f", (&Cal).Addition()))
	fmt.Println("Num1 - Num2 = ", (&Cal).Sub())
	fmt.Println("Num1 * Num2 = ", (&Cal).Multi())
	fmt.Println("Num1 / Num2 = ", (&Cal).Devid())
	fmt.Println("Num1 / Num2 = ", (&Cal).gerRes('/'))
}

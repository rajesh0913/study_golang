package main

// import "fmt"
// import "unsafe"
import (
	"fmt"
)

func main() {
	//演示小数类型使用
	var price float32 = 89.12
	fmt.Println("price = ", price)
	//浮点数
	var num1 float32 = -0.00000089
	var num2 float64 = -789234.00000089
	fmt.Println("num1 = ", num1, "num2 = ", num2)
	//精度丢失
	var num3 float32 = 13.00000901
	fmt.Println("num3 = ", num3)
	//声明默认值
	var num5 = 1.1
	fmt.Printf("num5的默认数据类型是：%T\n", num5)
	//科学计数法
	num6 := 1.234e3
	num7 := 1.234e-3
	fmt.Println("num6 = ", num6, "num7 = ", num7)

}

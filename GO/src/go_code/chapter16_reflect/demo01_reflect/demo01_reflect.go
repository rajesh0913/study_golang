package main

import (
	"fmt"
	"reflect"
)

/*
	1.反射可以动态获取变量的各种信息
	2.如果是结构体变量，还可以获取结构体本身信息
	3.通过反射，可以修改变量的值，还可以调用关联的方法
	4.反射包：reflect

	reflect.TypeOf:获取变量类型，返回reflect.Type类型const
	reflect.ValueOf：获取变量的值，反馈reflect.Value类型
*/

//1. 通过反射获取传入变量的type、kind、value
func reflectTest01(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println(rType)

	//valueOf
	rVal := reflect.ValueOf(b) //reflect.value类型
	fmt.Printf("rval = %v, rval type = %T\n", rVal, rVal)
	n1 := int64(10)
	n2 := n1 + rVal.Int() //转换类型
	fmt.Println(n2)

	//转回interface{}
	iv := rVal.Interface()
	//类型断言转成需要的类型
	num2 := iv.(int)
	fmt.Println("num2 = ", num2)

}

//2.对结构体的反射
func reflectTest02(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println(rType)

	//valueOf
	rVal := reflect.ValueOf(b) //reflect.value类型
	fmt.Printf("rval = %v, rval type = %T\n", rVal, rVal)

	//转回interface{}
	iv := rVal.Interface()
	fmt.Printf("iv = %v  iv type = %T\n", iv, iv)
	//可以使用switch断言形式做的更加灵活
	stu, ok := iv.(Student)
	if ok {
		fmt.Printf("stu.Name = %v\n", stu.Name)
	}

	//获取kind
	fmt.Println(rVal.Kind())
	fmt.Println(rType.Kind())
}

type Student struct {
	Name string
	Age  int
}

func main() {

	//1
	var num int = 100
	reflectTest01(num)

	//2
	stu := Student{
		"Tom",
		20,
	}
	reflectTest02(stu)
}

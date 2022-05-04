package main

import (
	"fmt"
	"reflect"
)

/*
	1.reflect.value.Kind获取变量类别，返回常量
	2.Type是类型，Kind是类别，Type和Kind可能是相同，也可能不同（结构体）
	3.通过反射可以让变量在interface和value之间相互转换
	4.使用反射方式获取变量的值，并返回对应的类型，要求数据类型匹配，如int要使用类型断言或者.Int()
	5.通过反射来修改变量，注意当使用setXxx方法设置需要通过对应的指针类型完成，同时需要使用reflect.Value.Elem()
*/

//5.
func reflect01(b interface{}) {
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal type= %v\n", rVal.Kind())

	//rVal.SetInt(20)->rVal是指针
	rVal.Elem().SetInt(20)

}

func main() {
	//5.通过反射修改值
	var num int = 100
	reflect01(&num)
	fmt.Println("num = ", num)
}

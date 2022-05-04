package main

import (
	"fmt"
	"reflect"
)

/*
	float64=1.2
	反射得到reflect.Value，获取对应Type、Kind、Value
	将reflect.Value转换成interface{}，再将interface{}转换成float

*/

func reflect01(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println("type = ", rType)
	rKind := rType.Kind()
	fmt.Println("kind = ", rKind)
	rVal := reflect.ValueOf(b)
	fmt.Println("value = ", rVal)

	iv := rVal.Interface()
	fmt.Printf("iv value = %v,iv type = %T\n", iv, iv)

	b2 := iv.(float64)
	fmt.Printf("b2 value = %v,b2 type = %T\n", b2, b2)
}

func main() {
	f1 := 1.2
	reflect01(f1)
}

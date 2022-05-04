package main

import (
	"fmt"
	"reflect"
)

/*
	使用反射遍历结构体字段，调用结构体方法，获取结构体标签值
*/
type Monster struct {
	Name  string `json:"monster_name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

func (s Monster) Print() {
	fmt.Println("---Start---")
	fmt.Println(s)
	fmt.Println("---end---")
}
func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}
func (s *Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}
func TestStruct(a interface{}) {
	//获取信息
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("except struct")
		return
	}

	//获取字段
	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num)
	for i := 0; i < num; i++ {
		fmt.Printf("fields %d: 值为%v\n", i, val.Field(i))
		//获取json
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("field %d: tag = %v\n", i, tagVal)
		}
	}
	//获取方法
	numOfmethod := val.NumMethod()
	fmt.Printf("struct has %d method\n", numOfmethod)
	//调用方法
	val.Method(1).Call(nil) //val.method()->第二个方法（按ASCII排序），call()调用

	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	//调用方法0
	res := val.Method(0).Call(params) //return []reflect.Value
	fmt.Println("res = ", res[0].Int())
}

func main() {
	var a Monster = Monster{
		Name:  "黄鼠狼精",
		Age:   400,
		Score: 30.8,
	}
	TestStruct(a)
}

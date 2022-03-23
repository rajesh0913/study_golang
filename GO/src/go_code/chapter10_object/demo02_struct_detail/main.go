package main

import (
	"encoding/json"
	"fmt"
)

type point struct {
	x int
	y int
}
type rect struct {
	leftUp, rightDown point
}

type A struct {
	Num int
}
type B struct {
	Num int
}

//3. 取别名,支持强制转换
type C A

//4. tag
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	//1. 结构体所有字段内存中连续
	r1 := rect{point{1, 2}, point{3, 4}}
	fmt.Printf("rect.leftUp.x address = %p\trect.leftUp.y address = %p\nrect.rightDown.x address = %p\trect.rightDown.y address = %p\n",
		&r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)

	//2. 结构体是单独类型，和其他类型转换所有必须一致
	var a A
	var b B
	var c C
	//A=B ->err
	//强制转换
	a = A(b)
	b = B(c)
	fmt.Println(a, b)

	//4. tag
	//用于序列化和反序列化
	monster := Monster{"牛魔王", 500, "芭蕉扇"}
	jsonmonster, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json err: ", err)
	} else {

		fmt.Println(string(jsonmonster))
	}

}

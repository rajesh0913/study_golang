package main

import "fmt"

//声明定义结构体，将字段属性放入结构中
type cats struct {
	Name  string
	Age   int
	Color string
}
type Person struct {
	Name  string
	Age   int
	Score [5]float64
	ptr   *int              //nil,没有分配空间
	slice []int             //nil
	map1  map[string]string //nil
}

func main() {
	//结构体变量
	var cat cats
	fmt.Println(cat)
	cat.Name = "小白"
	cat.Age = 4
	cat.Color = "white"
	fmt.Println(cat)

	//方式1：声明变量
	var P1 Person
	//方式2：
	//P2 := Person{}
	P2 := Person{}
	//方式3
	//P3为指针，赋值方式与前两种不一样
	var P3 *Person = new(Person)
	//方式4
	//P4为指针
	var P4 *Person = &Person{}

	(*P3).Name = "papy"
	P3.Age = 33
	fmt.Println(P1, "\n", P2, "\n", *P3, "\n", P4)
	//ptr、slice、map都没有分配空间
	if P1.ptr == nil {
		fmt.Println("ook1")
	}
	if P1.slice == nil {
		fmt.Println("ook2")
	}
	if P1.map1 == nil {
		fmt.Println("ook3")
	}
	//**********分配空间才能使用**************//
	P1.slice = make([]int, 10)
	P1.slice[0] = 100
	fmt.Println(P1.slice)

	P1.map1 = make(map[string]string)
	P1.map1["cat"] = "tom"
	P1.map1["mouse"] = "jerry"
	fmt.Println(P1.map1)

	//结构体是值类型，默认值拷贝，互不影响

}

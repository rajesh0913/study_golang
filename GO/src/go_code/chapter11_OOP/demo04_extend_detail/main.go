package main

import "fmt"

type A struct {
	Name string
	age  int
}

func (a *A) SayOK() {
	fmt.Println("A say ok", a.Name)
}
func (a *A) hello() {
	fmt.Println("A hello", a.Name)
}

type B struct {
	A
}

//5. 组合(嵌套有名结构体),可以嵌套多个
type C struct {
	a A
}

func main() {
	var b B
	b.A.Name = "tom"
	//1. 匿名结构体的隐藏字段可以被继承和访问
	b.A.age = 19
	b.A.SayOK()
	b.A.hello()
	//2. 可以简化(不声明匿名结构体)
	b.age = 22
	//3. 查找就近原则
	//4. 包含两个匿名且有相同字段的结构体，访问必须指明
	//5. 访问有名结构体
	var c C
	c.a.age = 18
	//6. 创建时指明各个匿名结构体的值
	var d = C{
		a: A{
			Name: "jack",
			age:  55,
		},
	}
	fmt.Println(d.a.age)
}

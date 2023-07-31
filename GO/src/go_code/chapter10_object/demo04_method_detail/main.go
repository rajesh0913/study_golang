package main

import "fmt"

type A struct {
	Num int
}

//2. 为提高效率(从值拷贝变为地址传递），通常方法和结构体的指针类型进行绑定
func (A *A) function() int {
	//fmt.Println((*A).Num)
	return (*A).Num * (*A).Num
}

type Student struct {
	Name string
	Age  int
}

//5. 重写，实现string()功能
func (stu *Student) String() string {
	str := fmt.Sprintf("Name = [%v],Age = [%v].", stu.Name, stu.Age)
	return str
}
func main() {
	//1. 值类型传递
	var a A
	//2.传入指针，则会修改原来的字段内容
	//a.Num = 22
	fmt.Println((&a).function())

	//3. 所有自定义类型（包括int float等）都可以有方法
	//4. 大小写首字母同样控制访问域
	//5. 为一个方法实现string()的功能，fmt.Println()会默认调用该函数
	Stu := Student{"qinzhibao", 25}
	fmt.Println("stu", Stu)
	fmt.Println("默认调用string():", &Stu)
	//6. 创建结构体变量，指定字段值
	b := Student{
		"qin", 18,
	}
	c := Student{
		Name: "zhi",
		Age:  20,
	}
	//返回结构体指针类型
	var d *Student = &Student{
		"bao", 22,
	}
	fmt.Println(b, c, d)
	e := 10
	f := &e
	fmt.Println(&e)
	fmt.Println(f)
}

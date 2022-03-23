package encapsulation

import "fmt"

//1. 将结构体、字段首字母小写进行封装
type person struct {
	Name   string
	age    int     //encapsulation
	salary float64 //encapsulation
}

//2. 给对应结构体提供工厂模式函数，首字母大写->构造函数
func NewPerson(name string) *person {
	return &person{
		Name:   name,
		age:    -1,
		salary: -1,
	}
	//return nil
}

//3. 提供Set()方法（pubulic），用于对属性进行判断和赋值
func (p *person) SetAge(age int) {

	if age <= 120 && age >= 0 {
		(*p).age = age
	} else {
		fmt.Println("年龄输入错误..")
	}
}
func (p *person) SetSalary(salary float64) {

	if salary >= 5000 && salary <= 30000 {
		(*p).salary = salary
	} else {
		fmt.Println("薪水输入错误..")
	}
}

//4. 提供Get()方法（pubulic），用于对属性进行判断和获取
func (p *person) GetAge() int {

	return (*p).age

}

func (p *person) GetSalary() float64 {

	return (*p).salary

}

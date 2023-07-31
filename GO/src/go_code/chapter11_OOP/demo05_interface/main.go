package main

import "fmt"

//interface可以声明一些方法，并且不需要实现
//不能包含任何变量
type Usb interface {
	//声明没有实现的方法
	Start() // 不能有方法体{}
	Stop()
}

//phone struct
type Phone struct {
}

//phone 实现interface方法
func (p Phone) Start() {
	fmt.Println("phone starts working...")
}
func (p Phone) Stop() {
	fmt.Println("phone stops working...")
}

//camera struct
type Camera struct {
}

//camera 实现interface方法
func (p Camera) Start() {
	fmt.Println("camera starts working...")
}
func (p Camera) Stop() {
	fmt.Println("camera stops working...")
}

//computer
type Computer struct {
}

//接收usb接口类型变量

func (c Computer) Working(usb Usb) {
	//通过接口变量调用接口方法
	usb.Start()
	usb.Stop()
}
func main() {
	c := Computer{}
	p := Phone{}
	//	Camera := Camera{}

	c.Working(p)

}

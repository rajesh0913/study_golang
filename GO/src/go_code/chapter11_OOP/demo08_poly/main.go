package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}

//phone struct
type Phone struct {
	Name string
}

//phone 实现interface方法
func (p Phone) Start() {
	fmt.Println("phone starts working...")
}
func (p Phone) Stop() {
	fmt.Println("phone stops working...")
}

//assert-phone 特有方法
func (p Phone) Call() {
	fmt.Println("phone is calling...")
}

//camera struct
type Camera struct {
	Name string
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
	//类型断言：phone单独调用call()
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}
func main() {
	//定义接口数组存放phone和camera结构体变量
	//不是接口不能存储
	var p Phone
	var c Computer
	//1. 多态参数
	c.Working(p)
	//2. 多态数组
	var usbArr [3]Usb
	usbArr[0] = Phone{"xiaomi"}
	usbArr[1] = Phone{"vivo"}
	usbArr[2] = Camera{"nikon"}
	fmt.Println(usbArr)
	for _, value := range usbArr {
		c.Working(value)
	}

}

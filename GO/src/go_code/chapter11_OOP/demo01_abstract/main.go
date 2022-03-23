package main

import "fmt"

//抽象
//结构体作为抽象对象
type Account struct {
	//字段存取抽象对象属性
	AccountNo string
	Pwd       string
	Balance   float64
}

//方法作为抽象对象可进行的操作
func (acc *Account) Deposite(money float64, pwd string) {
	//check in
	if pwd != (*acc).Pwd {
		fmt.Println("输入密码不正确！")
		return
	} else if money <= 0 {
		fmt.Println("输入金额错误！")
		return
	}
	(*acc).Balance += money
	fmt.Println("存款成功！")
	(*acc).Query(pwd)
}
func (acc *Account) WithDraw(money float64, pwd string) {
	//check in
	if pwd != (*acc).Pwd {
		fmt.Println("输入密码不正确！")
		return
	} else if money <= 0 || money > (*acc).Balance {
		fmt.Println("输入金额错误！")
		return
	}
	(*acc).Balance -= money
	fmt.Println("取款成功！")
	(*acc).Query(pwd)
}
func (acc *Account) Query(pwd string) {
	//check in
	if pwd != (*acc).Pwd {
		fmt.Println("输入密码不正确！")
		return
	}
	fmt.Println((*acc).AccountNo, "Balance = ", (*acc).Balance)
}
func main() {

	account := Account{
		"qinzhibao",
		"123456",
		9999,
	}
	(&account).Query("123456")
	(&account).WithDraw(1000, "123456")

}

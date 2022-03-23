package main

import (
	"fmt"
	"go_code/chapter12_program/demo02_customers_manage/model"
	"go_code/chapter12_program/demo02_customers_manage/service"
)

type Customer_View struct {
	key              string
	loop             bool
	Customer_Service *service.Customer_Service
}

func (cv *Customer_View) List() {
	list := cv.Customer_Service.List()
	fmt.Println("-------------------------客户列表-------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱\t")
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i].Get_info())
	}
	fmt.Printf("-------------------------客户列表完成-------------------------\n\n")
}

//get customer's information and add to list
func (cv *Customer_View) Add() {
	name := ""
	gender := ""
	age := 0
	phone := ""
	email := ""
	fmt.Println("-------------------------添加客户-------------------------")
	fmt.Println("姓名：")
	fmt.Scanln(&name)
	fmt.Println("性别：")
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	fmt.Scanln(&age)
	fmt.Println("电话：")
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	fmt.Scanln(&email)

	customer := model.New_Custonmer_no_id(name, gender, age, phone, email)

	if (*cv).Customer_Service.Add(customer) {
		fmt.Println("-------------------------添加完成-------------------------")
	} else {
		fmt.Println("-------------------------添加失败-------------------------")
	}

}

func (cv *Customer_View) Delete() {
	id := -1
	fmt.Println("-------------------------删除客户-------------------------")
	fmt.Println("输入编号：")
	fmt.Scanln(&id)
	if id == -1 {
		return
	} else {
		fmt.Println("确认删除？（Y/N）")
		choice := ""
		fmt.Scanln(&choice)
		if choice == "Y" || choice == "y" {
			if (*cv).Customer_Service.Delete(id) {
				fmt.Println("-------------------------删除完成-------------------------")
			} else {
				fmt.Println("-------------------------删除失败，输入id号不存在-------------------------")
			}

		}
	}

}

func (cv *Customer_View) main_menu() {
	for {
		fmt.Println("-------------------------客户信息管理软件-------------------------")
		fmt.Println("                         1. 添 加 客 户")
		fmt.Println("                         2. 修 改 客 户")
		fmt.Println("                         3. 删 除 客 户")
		fmt.Println("                         4. 客 户 列 表")
		fmt.Println("                         5. 退 出")
		fmt.Print("请选择(1-5):")
		fmt.Scanln(&(*cv).key)

		switch (*cv).key {
		case "1":
			fmt.Println("1. 添 加 客 户")
			(*cv).Add()
		case "2":
			fmt.Println("2. 修 改 客 户")
		case "3":
			fmt.Println("3. 删 除 客 户")
			(*cv).Delete()
		case "4":
			cv.List()
		case "5":
			(*cv).loop = false
		default:
			fmt.Println("输入有误，请重新输入(1-5)...")

		}
		if !(*cv).loop {
			break
		}
	}
	fmt.Println("退出客户管理系统...")
}
func main() {
	customer_view := Customer_View{
		key:              "",
		loop:             true,
		Customer_Service: service.New_Customer_Service(),
	}

	customer_view.main_menu()
}

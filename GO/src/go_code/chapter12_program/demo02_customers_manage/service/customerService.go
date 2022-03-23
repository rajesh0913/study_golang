package service

import (
	"go_code/chapter12_program/demo02_customers_manage/model"
)

type Customer_Service struct {
	customer     []model.Customer
	customer_Num int
}

func New_Customer_Service() *Customer_Service {
	customer_Service := &Customer_Service{}
	customer_Num := 1
	customer := model.New_Custonmer(1, "张三", "男", 20, "112", "zs@gmail.com")
	customer_Service.customer = append(customer_Service.customer, customer)
	customer_Service.customer_Num = customer_Num
	return customer_Service
}

func (cv *Customer_Service) List() []model.Customer {
	return (*cv).customer
}
func (cv *Customer_Service) Add(customer model.Customer) bool {
	(*cv).customer_Num++
	customer.Id = (*cv).customer_Num
	(*cv).customer = append((*cv).customer, customer)
	return true
}
func (cv *Customer_Service) Delete(id int) bool {
	index := (*cv).Fide_Id(id)
	if index == -1 {
		return false
	}
	(*cv).customer = append((*cv).customer[:index], (*cv).customer[index+1:]...) //两个切片合成一个
	return true
}
func (cv *Customer_Service) Fide_Id(id int) int {
	index := -1
	for _, value := range (*cv).customer {
		if value.Id == id {
			index = id
			return index
		}
	}
	return index

}

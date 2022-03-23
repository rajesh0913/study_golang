package model

import "fmt"

type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

func New_Custonmer(Id int, Name string, Gender string, Age int, Phone string, Email string) Customer {
	return Customer{
		Id:     Id,
		Name:   Name,
		Gender: Gender,
		Age:    Age,
		Phone:  Phone,
		Email:  Email,
	}

}
func New_Custonmer_no_id(Name string, Gender string, Age int, Phone string, Email string) Customer {
	return Customer{
		Name:   Name,
		Gender: Gender,
		Age:    Age,
		Phone:  Phone,
		Email:  Email,
	}

}

func (cv Customer) Get_info() string {
	return fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", cv.Id, cv.Name, cv.Gender, cv.Age, cv.Phone, cv.Email)
}

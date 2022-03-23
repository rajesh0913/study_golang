package main

import (
	"fmt"
)

type Student struct {
	Name   string  `json:"name"`
	Gender string  `json:"gender"`
	Age    int     `json:"age"`
	Id     int     `json:"id"`
	Score  float64 `json:"scord"`
}

func (stu *Student) say() string {
	infoStr := fmt.Sprintf("student的信息:\n name = %v\tgender = %v\tage = %v\tid = %v\tscore = %v",
		(*stu).Name, (*stu).Gender, (*stu).Age, (*stu).Id, (*stu).Score)
	return infoStr
}

func main() {
	Stu := Student{
		Name:   "tom",
		Gender: "male",
		Age:    18,
		Id:     201410213055,
		Score:  99.9,
	}
	fmt.Println((&Stu).say())
}

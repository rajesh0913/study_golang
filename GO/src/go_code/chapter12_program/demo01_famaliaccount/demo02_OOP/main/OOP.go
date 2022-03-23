package main

import (
	"go_code/chapter12_program/demo01_famaliaccount/demo02_OOP/utils"
)

func main() {
	var Qin *utils.Famalit_Account = utils.New_Famaliy_Account()
	//fmt.Println(Qin)
	(*Qin).Show_Menu()
}

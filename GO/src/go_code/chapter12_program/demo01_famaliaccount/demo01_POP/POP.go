package main

import "fmt"

func main() {
	//接收选项
	var key string
	//默认余额
	balance := 0.0
	//收支金额
	money := 0.0
	//收支说明
	note := ""
	//记录收支明细
	details := "收支\t账户金额\t收支金额\t说明\n"
	//是否有收支行为
	isCount := false
	//是否循环
	var loop bool = true

	for {
		fmt.Println("-------------------------家庭收支记账软件-------------------------")
		fmt.Println("                         1. 收支明细")
		fmt.Println("                         2. 登记收入")
		fmt.Println("                         3. 登记支出")
		fmt.Println("                         4. 退出软件")
		fmt.Print("请选择(1-4):")
		fmt.Scanln(&key)
		switch key {
		case "1":
			if isCount {
				fmt.Println("-------------------------当前收支明细记录-------------------------")
				fmt.Println(details)
			} else {
				fmt.Println("当前无任何收支明细，请来一笔吧！")
			}
		case "2":
			fmt.Print("本次收入金额：")
			fmt.Scanln(&money)
			balance += money
			fmt.Print("本次收入说明：")
			fmt.Scanln(&note)
			details += fmt.Sprintf("收入\t%v\t\t%v\t\t%v\n", balance, money, note)
		case "3":
			fmt.Println("登记支出：")
			fmt.Print("本次支出金额：")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("登记失败！余额不足...")
				break
			}
			balance -= money
			fmt.Print("本次支出说明：")
			fmt.Scanln(&note)
			details += fmt.Sprintf("支出\t%v\t\t%v\t\t%v\n", balance, money, note)
		case "4":
			fmt.Printf("确认退出(Y/N):")
			var A string
			for {
				fmt.Scanln(&A)
				if A == "y" || A == "Y" {
					loop = false
					break
				} else if A == "n" || A == "N" {
					break
				} else {
					fmt.Println("输入错误！请重新输入(Y/N):")
				}

			}
		default:
			fmt.Println("输入有误！请输入正确选项！")
		}
		if !loop {
			break
		}
	}
	fmt.Println("已退出记账软件！")
}

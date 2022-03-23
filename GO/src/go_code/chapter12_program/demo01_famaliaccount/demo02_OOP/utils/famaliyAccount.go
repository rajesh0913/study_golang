package utils

import "fmt"

type Famalit_Account struct {
	//接收选项
	key string
	//默认余额
	balance float64
	//收支金额
	money float64
	//收支说明
	note string
	//记录收支明细
	details string
	//是否有收支行为
	isCount bool
	//是否退出
	loop bool
}

func (t *Famalit_Account) Show_Detail() {
	if (*t).isCount {
		fmt.Println("-------------------------当前收支明细记录-------------------------")
		fmt.Println((*t).details)
	} else {
		fmt.Println("当前无任何收支明细，请来一笔吧！")
	}
}

func (t *Famalit_Account) Income() {
	fmt.Print("本次收入金额：")
	fmt.Scanln(&(*t).money)
	(*t).isCount = true
	(*t).balance += (*t).money
	fmt.Print("本次收入说明：")
	fmt.Scanln(&(*t).note)
	(*t).details += fmt.Sprintf("收入\t%v\t\t%v\t\t%v\n", (*t).balance, (*t).money, (*t).note)
}

func (t *Famalit_Account) Outcome() {
	fmt.Println("登记支出：")
	fmt.Print("本次支出金额：")
	fmt.Scanln(&(*t).money)
	if (*t).money > (*t).balance {
		fmt.Println("登记失败！余额不足...")
		return
	}
	(*t).balance -= (*t).money
	(*t).isCount = true
	fmt.Print("本次支出说明：")
	fmt.Scanln(&(*t).note)
	(*t).details += fmt.Sprintf("支出\t%v\t\t%v\t\t%v\n", (*t).balance, (*t).money, (*t).note)
}

func (t *Famalit_Account) Exit() {
	fmt.Printf("确认退出(Y/N):")
	var A string
	for {
		fmt.Scanln(&A)
		if A == "y" || A == "Y" {
			(*t).loop = false
			break
		} else if A == "n" || A == "N" {
			break
		} else {
			fmt.Println("输入错误！请重新输入(Y/N):")
		}

	}
}

func (t *Famalit_Account) Show_Menu() {
	for {
		fmt.Println("-------------------------家庭收支记账软件-------------------------")
		fmt.Println("                         1. 收支明细")
		fmt.Println("                         2. 登记收入")
		fmt.Println("                         3. 登记支出")
		fmt.Println("                         4. 退出软件")
		fmt.Print("请选择(1-4):")
		fmt.Scanln(&(*t).key)
		switch (*t).key {
		case "1":
			(*t).Show_Detail()
		case "2":
			(*t).Income()
		case "3":
			(*t).Outcome()
		case "4":
			(*t).Exit()
		default:
			fmt.Println("输入有误！请输入正确选项！")
		}
		if !(*t).loop {
			break
		}
	}
	fmt.Println("已退出记账软件！")
}

//工厂模式
func New_Famaliy_Account() *Famalit_Account {
	return &Famalit_Account{
		//接收选项
		key: "",
		//默认余额
		balance: 0.0,
		//收支金额
		money: 0.0,
		//收支说明
		note: "",
		//记录收支明细
		details: "收支\t账户金额\t收支金额\t说明\n",
		//是否有收支行为
		isCount: false,
		//是否循环
		loop: true,
	}
}

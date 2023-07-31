package main

import "fmt"

//这个示例跟要讲解的不符合
type Monkey struct {
	Name string
}
type Flying interface {
	flying()
}

func (m *Monkey) climbing() {
	fmt.Println("爬树")
}

type LittleMonkey struct {
	Monkey //继承
}

func (t *LittleMonkey) Flying() {
	fmt.Println("学会飞翔")
}

func main() {
	monkey := LittleMonkey{
		Monkey{"wukong"},
	}
	monkey.climbing()
	(&monkey).Flying()

}

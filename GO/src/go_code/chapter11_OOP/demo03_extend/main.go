package main

import "fmt"

//共有属性
type Student struct {
	Name  string
	Age   int
	Score int
}

//共有方法
func (p *Student) ShowInfo() {
	fmt.Printf("Name = %v, Age = %v, Score = %v.", (*p).Name, (*p).Age, (*p).Score)
}

func (p *Student) SetScore(score int) {
	(*p).Score = score
}

//小学生考试系统
type Pupil struct {
	Student //嵌入继承的匿名结构体
}

// func (p *Pupil) ShowInfo() {
// 	fmt.Printf("Name = %v, Age = %v, Score = %v.", (*p).Name, (*p).Age, (*p).Score)
// }

// func (p *Pupil) SetScore(score int) {
// 	(*p).Score = score
// }

//私有属性
func (p *Pupil) Testing() {
	fmt.Println("小学生考试中...")
}

type Graduate struct {
	Student //嵌入继承的匿名结构体
}

//私有属性
func (p *Graduate) Testing() {
	fmt.Println("大学生考试中...")
}

func main() {
	var pupil = &Pupil{
		Student{
			Name: "tom",
			Age:  10,
		},
	}
	pupil.Student.Age = 11
	(*pupil).Testing()
	(*pupil).Student.SetScore(98)
	(*pupil).Student.ShowInfo()

	var graduate = &Graduate{
		Student{
			"jack",
			22,
			0,
		},
	}
	graduate.Student.Age = 23
	(*graduate).Testing()
	(*graduate).Student.SetScore(98)
	(*graduate).Student.ShowInfo()

}

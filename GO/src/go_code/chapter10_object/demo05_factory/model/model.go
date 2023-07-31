package model

//小写字母开头，不能直接使用
//跨包使用->工厂模式
type student struct {
	Name  string
	Age   int
	score float64
}

func NewStudent(n string, s int) *student {
	return &student{
		Name: n,
		Age:  s,
	}
}

//score不可以直接调用，提供方法返回对应参数
func (stu *student) GetScore() float64 {
	return stu.score
}

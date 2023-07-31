package main

import "fmt"

type A struct {
	Name string
}

//绑定method
//值传递
func (a A) test() {
	a.Name = "jack"
	fmt.Println("test()", a.Name)
}
func (a A) count(n int) (sum int) {
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum
}

type B struct {
}

func main() {
	var a A
	//var b B
	a.Name = "tom"
	a.test()
	fmt.Println(a.Name)
	fmt.Println(a.count(10))

	var b A
	b.Name = "jerry"
	b.test()
	fmt.Println(b.Name)

}

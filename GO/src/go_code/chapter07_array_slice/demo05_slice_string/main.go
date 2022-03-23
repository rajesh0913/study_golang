package main

import "fmt"

//string底层是一个byte数组，可以进行切片处理
func main() {
	str := "qinzhibao0817@gmail.com"

	slice := str[14:]
	fmt.Println(slice)
	//字符串需要修改：string -> []byte ->string. []byte不能处理中文，将[]byte改为[]rune
}

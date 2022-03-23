package main

import (
	"fmt"
)

func main() {
	//演示string类型使用
	var address string = "天秀花园"
	fmt.Print(address)
	//golang中：字符串一旦赋值，字符串不可修改

	//两种表现形式：
	//形式1：双引号，可识别转义字符
	str1 := "abc\nabc"
	fmt.Println(str1)
	//形式2：反引号，以字符串的原生形式输出，包括换行和特殊字符，可以防止攻击
	//输出源代码的效果
	str2 := `//演示string类型使用
	var address string = "天秀花园"
	fmt.Print(address)
	//golang中：字符串一旦赋值，字符串不可修改

	//两种表现形式：
	//形式1：双引号，可识别转义字符
	str1:="abc\nabc"
	fmt.Println(str1)
	//形式2：反引号，以字符串的原生形式输出，包括换行和特殊字符，可以防止攻击
	//输出源代码的效果
	`
	fmt.Println(str2)
	//字符串拼接方式
	var str3 = "hello " + "world"
	str3 += " hahaha!"
	fmt.Println(str3)
	//拼接很多字符串，加号应该放在上面一行
}

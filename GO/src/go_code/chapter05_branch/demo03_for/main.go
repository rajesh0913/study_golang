package main

import (
	"fmt"
)

func main() {
	//case1
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	fmt.Print("\n")
	//case2
	j := 0
	for j < 10 {
		fmt.Print(j)
		j++
	}
	fmt.Print("\n")
	//case3:while方式
	k := 0
	for {
		if k < 10 {
			fmt.Print("ok")
		} else {
			break
		}
		k++
	}
	//case4:do-while方式
	m := 0
	for {
		fmt.Print("ok")
		m++
		if m >= 10 {
			break
		}
	}
	//遍历字符串1:传统方式,传统方式对字符串使用【字节】进行遍历
	//但是【中文】utf8对应3个字节，需要将str转成切片
	str := "hello world"
	str2 := "hello world 北京"
	str3 := []rune(str2)
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}
	for i := 0; i < len(str3); i++ {
		fmt.Printf("%c\n", str3[i])
	}
	//遍历字符串2:for-range方式  按照【字符】进行遍历
	for index, val := range str2 {
		fmt.Printf("%d\t%c\n", index, val)
	}
}

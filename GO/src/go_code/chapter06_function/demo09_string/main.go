package main

import (
	"fmt"
	"strconv"
	"strings"
)

//字符串本身不可修改！
func main() {

	//1.len(v type) int :统计长度，根据参数类型
	//golang统一使用utf-8，ascii码占用一个字节，中文占用三个字节
	str1 := "hello北"
	fmt.Println("str1 的长度为", len(str1)) //5+3=8

	//2.字符串遍历时，同时处理有中文的问题 r:=[]rune(str)
	str2 := "hello北京"
	str3 := []rune(str2)
	for i := 0; i < len(str3); i++ {
		fmt.Printf("str3 字符=%c\n", str3[i])
	}

	//3.字符串转整数
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误")
	} else {
		fmt.Println("转换结果为：", n)
	}

	//4.整数转换成字符串
	str4 := strconv.Itoa(12345)
	fmt.Println("str4 = ", str4)

	//5.字符串转换成byte
	bytes := []byte("hello go")
	fmt.Printf("bytes = %v\n", bytes)

	//6.byte转换成字符串
	str5 := string([]byte{97, 98, 99})
	fmt.Println("str5 = ", str5)

	//7.10进制转成2，8，16进制，返回字符串
	str6 := strconv.FormatInt(123, 2)
	fmt.Println("123 对应的2进制数为： ", str6)

	str6 = strconv.FormatInt(123, 8)
	fmt.Println("123 对应的8进制数为： ", str6)

	str6 = strconv.FormatInt(123, 16)
	fmt.Println("123 对应的16进制数为： ", str6)

	//8.查找子串是否在指定字符串【重点】
	isContain := strings.Contains("hello world", "hello")
	if isContain {
		fmt.Println("Yes!")
	} else {
		fmt.Println("No!")
	}

	//9.统计字符串中有几个指定的字串
	num := strings.Count("hello world", "l")
	fmt.Println("num = ", num)

	//10.不区分大小写进行字符串比较
	fmt.Println("不区分大小写进行比较：", strings.EqualFold("abc", "ABC"))
	fmt.Println("区分大小写进行比较：", "abc" == "ABC")

	//11.返回子串第一次出现在字符串中的index位置
	fmt.Println("返回子串第一次出现在字符串中index的位置：", strings.Index("qin_abc", "abc"))

	//12.返回子串最后一次出现在字符串中的index位置
	fmt.Println("返回子串最后一次出现在字符串中index的位置：", strings.LastIndex("qin_abc_abc_abc", "abc"))

	//13.将子串替换成另一个子串,返回新字符串（未改变旧字符串）
	str7 := "qin_abc_abc_abc"
	str8 := strings.Replace("qin_abc_abc_abc", "abc", "gogogo", 2)
	fmt.Printf("将str7: %v子串替换成另一个子串\tstr8：%v\n", str7, str8)

	//14.按照指定字符，分割字符串为字符串数组
	strArr9 := strings.Split("hello,world", "l")
	fmt.Println("分割字符串为字符串数组\n", strArr9)
	fmt.Println(len(strArr9))
	for i := 0; i < len(strArr9); i++ {
		fmt.Printf("%v\n", strArr9[i])
	}

	//15.字符串进行大小写字母转换
	fmt.Println("字符串进行大小写字母转换", strings.ToUpper("qin_abc_abc_abc"))
	fmt.Println("字符串进行大小写字母转换", strings.ToLower("QIN_ABC_ABC_ABC"))

	//16.消除字符串左右两边空格
	str11 := strings.TrimSpace("     消除字符串左右两边空格     ")
	fmt.Println("消除字符串左右两边空格\t=", str11)
	str11_strconv := []rune(str11)
	for i := 0; i < len(str11_strconv); i++ {
		fmt.Printf("%q", str11_strconv[i])
	}
	fmt.Println()

	//17.消除字符串左右两边指定字符
	str10 := strings.Trim("  !消除字符串左右两边指定字符    !", " !")
	str10_strconv := []rune(str10)
	for i := 0; i < len(str10_strconv); i++ {
		fmt.Printf("%q", str10_strconv[i])
	}
	fmt.Println()

	//18.消除字符串左边指定字符
	str12 := strings.TrimLeft("     !消除字符串左边指定字符    !      ", " !")
	fmt.Println(str12)
	//19.消除字符串右边指定字符
	str13 := strings.TrimRight("     !消除字符串右边指定字符    !      ", " !")
	fmt.Println(str13)

	//20.判断字符串是否以指定子串开头
	str14 := "ftp:\\192.168.0.1"
	str15 := "ftp"
	fmt.Printf("判断字符串%q是否以指定子串%q开头: %v\n", str14, str15, strings.HasPrefix(str14, str15))

	//21.判断字符串是否以指定子串结尾
	fmt.Printf("判断字符串%q是否以指定子串%q结尾: %v\n", str14, str15, strings.HasSuffix(str14, str15))

}

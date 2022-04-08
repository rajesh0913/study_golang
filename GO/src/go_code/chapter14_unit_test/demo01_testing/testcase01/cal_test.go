package cal

import (
	"testing"
)

/*
	1.测试文件名必须以_test.go结尾
	2.测试用例函数必须以Test开头，一般为：Test+被测试函数名
	3.Testfunction(t *testing.T)形参类型必须是*testing.T
	4.一个测试文件可以有多个测试用例函数
	5.运行测试用例指令、
		1)cmd>go test		正确无日志，错误输出日志
		2)cmd>go test -v	输出日志
	6.出现错误，使用t.Fatalf格式化输出错误信息，并退出程序
	7.t.Logf方法可以输出相应日志
	8.测试用例函数并没有放在main中，也执行了->测试用例方便之处
	9.PASS表示测试用例运行成功，FAIL表示失败
	10.测试单个文件，一定要带上被测试的源文件(不带则会测试全部包含的测试方法)
		go test -v cal_test.go cal.go
	11.测试单个方法
		go test -v -run=TestAddUpper



*/

//编写要给测试用例，检测是否正确
func TestAddUpper(t *testing.T) { //Test后接第一个字符不能是（a-z）
	//调用
	res := addUpper(10)
	if res != 55 {
		//fmt.Printf("addUpper(10)错误，返回值： %v   期望值： %v", res, 55)
		t.Fatalf("addUpper(10)错误，返回值： %v   期望值： %v", res, 55)
	}
	//正确，进入日志
	t.Logf("addUper(10)正确执行...")

}

//编写要给测试用例，检测是否正确
func TestGetsub(t *testing.T) { //Test后接第一个字符不能是（a-z）
	//调用
	res := getsub(10, 3)
	if res != 7 {
		//fmt.Printf("getsub(10,3)错误，返回值： %v   期望值： %v", res, 7)
		t.Fatalf("getsub(10,3)错误，返回值： %v   期望值： %v", res, 7)
	}
	//正确，进入日志
	t.Logf("getsub(10,3)正确执行...")

}

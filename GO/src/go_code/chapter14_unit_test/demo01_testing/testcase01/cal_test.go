package cal

import (
	"testing"
)

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

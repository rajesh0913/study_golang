package main

// import "fmt"
// import "unsafe"
import (
	"fmt"
	"unsafe"
)

func main() {
	//测试变量范围
	var i int8 = 127
	fmt.Println("i = ", i)
	//整型使用细节
	var n1 = 100 //n1是什么类型？
	fmt.Printf("n1的数据类型 %T \n", n1)
	//查看程序中某个变量的占用字节大小和数据类型（重点
	var n2 int64 = 10
	//unsafe.sizeof()是unsafe包中函数，可以返回变量字节数量
	fmt.Printf("n2 的类型 %T\t\tn2 占用字节数 %d", n2, unsafe.Sizeof(n2))

}

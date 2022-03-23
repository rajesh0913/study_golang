package main

import "fmt"

//判断输入数据类型
func TypeJudge(data ...interface{}) {
	for index, value := range data {
		switch value.(type) {
		case bool:
			fmt.Printf("data[%v] type = bool,value = %v.\n", index, value)
		case float32:
			fmt.Printf("data[%v] type = float32,value = %v.\n", index, value)
		case float64:
			fmt.Printf("data[%v] type = float64,value = %v.\n", index, value)
		case int, int32, int64:
			fmt.Printf("data[%v] type = int,value = %v.\n", index, value)
		case string:
			fmt.Printf("data[%v] type = string,value = %v.\n", index, value)
		default:
			fmt.Printf("data[%v] type = unknown,value = %v.\n", index, value)
		}
	}
}
func main() {
	var v1 float32 = 1.1
	var v2 float64 = 2.20
	var v3 int = 10
	var v4 string = "123456"
	var v5 byte = 'Q'
	TypeJudge(v1, v2, v3, v4, v5)
}

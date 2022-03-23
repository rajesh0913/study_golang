package main

import (
	"fmt"
	"os"
)

func main() {
	//1. file.open
	//file对象
	//file指针
	//file文件句柄
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("open file err: ", err)
	}
	fmt.Printf("%v", file) //指针类型

	//2. file.close
	err = file.Close()
	if err != nil {
		fmt.Println("close file err:a ", err)
	}

}

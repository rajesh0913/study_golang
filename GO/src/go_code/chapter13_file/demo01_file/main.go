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
	file, err := os.Open("chapter13_file/demo01_file/test.txt")
	if err != nil {
		fmt.Println("open file err: ", err)
	} else {
		fmt.Println("open file success!")
	}
	fmt.Printf("%v\n", file) //指针类型

	//2. file.close
	err = file.Close()
	if err != nil {
		fmt.Println("close file err:a ", err)
	} else {
		fmt.Println("close file success!")
	}

}

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("chapter13_file/demo02.file_reader/test.txt")
	if err != nil {
		fmt.Println("open file err: ", err)
	}
	//函数退出时，及时关闭file，防止内存泄漏
	defer file.Close() //先被 defer 的语句最后被执行，最后被 defer 的语句，最先被执行。

	//1. 方式1：带缓冲（适用于大文件
	reader := bufio.NewReader(file) //4096缓冲
	for {                           //重复读取直到结束
		str, err := reader.ReadString('\n') //换行结束一次读取
		if err == io.EOF {                  //文件末尾
			break
		}
		fmt.Print(str) //输出内容
	}
	fmt.Println("文件读取结束...")

	//2. 方式2：不带缓冲一次性读取（仅适用于小文件
	content, err2 := ioutil.ReadFile(file.Name()) //没有打开也没有关闭文件：被隐藏到函数中
	if err2 != nil {
		fmt.Println("open file err: ", err2)
	}
	fmt.Printf("%v", string(content))

}

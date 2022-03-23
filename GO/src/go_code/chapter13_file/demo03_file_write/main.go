package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	/**************************************************************/
	//1. 新建文件并写入内容
	file_path := "F:/GO/project/src/go_code/chapter13_file/demo03_file_write/test01.txt"
	file, err := os.OpenFile(file_path, os.O_WRONLY|os.O_CREATE, 0666) //os.O_CREATE
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}
	defer file.Close()

	str := "test01...\r\n" //不同编辑器识别不同换行符
	//写入时，使用带缓存的“*writer”
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	//writer带缓存，调用该方法，内容写入缓存，需要调用flush，将数据写到文件中
	writer.Flush()

	/**************************************************************/
	//2. 打开已经存在的文件，覆盖原有内容
	file_path = "F:/GO/project/src/go_code/chapter13_file/demo03_file_write/test02.txt"
	file, err = os.OpenFile(file_path, os.O_WRONLY|os.O_TRUNC, 0666) //os.O_TRUNC
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}

	defer file.Close()
	str = "test02...\r\n"
	//写入时，使用带缓存的“*writer”
	writer = bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	//writer带缓存，调用该方法，内容写入缓存，需要调用flush，将数据写到文件中
	writer.Flush()

	/**************************************************************/
	//3. 打开已经存在的文件，追加内容
	file_path = "F:/GO/project/src/go_code/chapter13_file/demo03_file_write/test03.txt"
	file, err = os.OpenFile(file_path, os.O_WRONLY|os.O_APPEND, 0666) //O_APPEND
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}

	defer file.Close()
	str = "test03...\r\n"
	//写入时，使用带缓存的“*writer”
	writer = bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	//writer带缓存，调用该方法，内容写入缓存，需要调用flush，将数据写到文件中
	writer.Flush()

	/**************************************************************/
	//4. 打开已经存在的文件，显示到终端，追加内容
	file_path = "F:/GO/project/src/go_code/chapter13_file/demo03_file_write/test04.txt"
	file, err = os.OpenFile(file_path, os.O_RDWR|os.O_APPEND, 0666) //O_APPEND
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}

	defer file.Close()

	//读取
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
	//写入时，使用带缓存的“*writer”
	str = "test04...\r\n"
	writer = bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	//writer带缓存，调用该方法，内容写入缓存，需要调用flush，将数据写到文件中
	writer.Flush()

}

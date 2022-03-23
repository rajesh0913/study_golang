package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//2. 判断是否存在
func Path_Exitsts(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { //返回nil，文件存在
		return true, nil
	}
	if os.IsNotExist(err) { //文件不存在
		return false, nil
	}
	return false, err //不确定是否存在
}

//3. 文件拷贝
func File_Copy(dir_file string, src_file string) (written int64, err error) {
	reader_file, err := os.Open(src_file)
	if err != nil {
		fmt.Printf("open file err: %v", err)
		return
	}
	defer reader_file.Close()

	reader := bufio.NewReader(reader_file)

	writer_file, err := os.OpenFile(dir_file, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err: %v", err)
		return
	}
	defer writer_file.Close()

	writer := bufio.NewWriter(writer_file)

	return io.Copy(writer, reader)

}

func main() {
	//1. 文件内容拷贝
	file_path01 := "F:/GO/project/src/go_code/chapter13_file/demo04_file_other/test01.txt"
	file_path02 := "F:/GO/project/src/go_code/chapter13_file/demo04_file_other/test02.txt"
	//读取到内存
	data, err := ioutil.ReadFile(file_path01)
	if err != nil {
		fmt.Println("read file01 err: ", err)
		return
	}
	err = ioutil.WriteFile(file_path02, data, 0666) //覆盖掉了
	if err != nil {
		fmt.Println("write file02 err: ", err)
		return
	}

	//3. 文件拷贝
	file_path01 = "F:/GO/project/src/go_code/chapter13_file/demo04_file_other/test_copy_reader/1.jpg"
	file_path02 = "F:/GO/project/src/go_code/chapter13_file/demo04_file_other/test_copy_writer/2.jpg"

	_, err = File_Copy(file_path02, file_path01)
	if err != nil {
		fmt.Println("拷贝错误,err: ", err)
	} else {
		fmt.Println("拷贝完成")
	}
}

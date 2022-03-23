package service

import (
	//"bufio"
	"fmt"
	"os"
	"strings"

	//"io"
	"io/ioutil"
)

func Get_file_address(address string, file string, store string) {
	var pathSeparator = string(os.PathSeparator)

	//读取新名文件内容（新名字）
	contents, err2 := ioutil.ReadFile(file) //没有打开也没有关闭文件：被隐藏到函数中

	if err2 != nil {
		fmt.Println("open file err: ", err2)
	}

	names := strings.Split(string(contents), "\r\n")
	//fmt.Println(names)

	files, _ := ioutil.ReadDir(address)
	fmt.Println()

	for index, file_in := range files {
		temp := address + pathSeparator
		//fmt.Println(index, len(names))
		//getCurrentDirectory()
		fmt.Println(file_in.Name())
		if index >= len(names) {
			temp = store + pathSeparator + fmt.Sprintf("%d", index+1) + ".mp3"
		} else {
			temp = store + pathSeparator + fmt.Sprintf("%d", index+1) + names[index] + ".mp3"
		}
		err := os.Rename(address+pathSeparator+file_in.Name(), temp)
		if err != nil {
			fmt.Println("open file err: ", err, "in", "index:", index, "old name:", file_in.Name())
		}

	}
}

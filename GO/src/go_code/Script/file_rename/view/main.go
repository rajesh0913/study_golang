package main

import (
	"fmt"
	"go_code/Script/file_rename/service"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
	var address string
	//var file string
	var store string
	fmt.Println("请输入文件夹地址：")
	fmt.Scanln(&address)
	fmt.Println("请输入存放文件地址(name.txt存放的文件夹)：")
	fmt.Scanln(&store)

	// exist, err := PathExists(store + string(os.PathSeparator) + "output")
	// if err != nil {
	// 	fmt.Printf("get dir error![%v]\n", err)
	// 	return
	// }

	// if exist {
	// 	fmt.Printf("has dir![%v]\n", store)
	// 	fmt.Println("clearing...")
	// 	os.Remove(store + string(os.PathSeparator) + "output")
	// 	os.Create(store + string(os.PathSeparator) + "output")
	// 	fmt.Println("clear done!")
	// } else {
	// 	fmt.Printf("no dir![%v]\n", store+string(os.PathSeparator)+"output")
	// 	// 创建文件夹
	// 	err := os.Mkdir(store+string(os.PathSeparator)+"output", os.ModePerm)
	// 	if err != nil {
	// 		fmt.Printf("mkdir failed![%v]\n", err)
	// 	} else {
	// 		fmt.Printf("mkdir success!\n")
	// 	}
	// }

	service.Get_file_address(address, store+string(os.PathSeparator)+"name.txt", address)
}

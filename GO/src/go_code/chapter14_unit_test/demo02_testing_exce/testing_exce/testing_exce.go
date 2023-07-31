package testing_exce

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

//绑定方法
func (this *Monster) Store() bool {
	//序列化
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("marshal err = ", err)
		return false
	}
	//保存文件

	filePath := "go_code/chapter14_unit_test/demo02_testing_exce/testing_exce/1.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("file open err = ", err)
		return false
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	//fmt.Println(string(data))
	_, err = writer.WriteString(string(data))
	writer.Flush()
	if err != nil {
		fmt.Println("write err = ", err)
		return false
	}
	return true
}
func (this *Monster) ReStore() bool {
	//文件读取序列化的字符串
	filePath := "go_code/chapter14_unit_test/demo02_testing_exce/testing_exce/1.txt"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("file read err = ", err)
		return false
	}
	//反序列化,data->[]byte
	fmt.Println("data = ", data)

	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Println("unmarshal err = ", err)
		return false
	}
	return true
}

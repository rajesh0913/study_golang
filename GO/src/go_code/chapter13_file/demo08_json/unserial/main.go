package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name   string  //`json:"monster_name"`
	Age    int     //`json:"monster_age"`
	Birth  string  //`json:"monster_birth"`
	Salary float64 //`json:"monster_salary"`
	Skill  string  //`json:"monster_skill"`
	//shadow bool    //`json:"monster_shadow"`
}

func unserialStruct() {
	str := "{\"Name\":\"牛魔王\",\"Age\":500,\"Birth\":\"2022.2.22\",\"Salary\":8000.0,\"Skill\":\"牛魔拳\"}"

	var monster Monster
	//fmt.Println(monster, []byte(str))
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Printf("unserial monster: %v", monster)
}

func unserialMap() {
	str := "{\"address\":\"洪崖洞\",\"age\":30,\"name\":\"红孩儿\"}"

	var a map[string]interface{} = make(map[string]interface{})
	//fmt.Println(monster, []byte(str))
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Printf("unserial a: %v", a)
}

func unserialSlice() {
	//通过程序获取而不通过原生数据，不需要转义处理
	str := "[{\"address\":\"北京\",\"age\":7,\"name\":\"jack\"},{\"address\":\"纽约\",\"age\":8,\"name\":\"tom\"}]"

	var slice []map[string]interface{}
	//make操作被封装到unmarshal函数中
	//var a map[string]interface{} = make(map[string]interface{})

	//fmt.Println(monster, []byte(str))
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Printf("unserial slice: %v", slice)
}

func main() {
	unserialStruct()
	fmt.Println()
	unserialMap()
	fmt.Println()
	unserialSlice()
}

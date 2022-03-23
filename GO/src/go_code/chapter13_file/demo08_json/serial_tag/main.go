package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name   string  `json:"monster_name"`
	Age    int     `json:"monster_age"`
	Birth  string  `json:"monster_birth"`
	Salary float64 `json:"monster_salary"`
	Skill  string  `json:"monster_skill"`
	shadow bool    //`json:"monster_shadow"`
}

func testStruct() {
	var monster Monster = Monster{
		Name:   "牛魔王",
		Age:    500,
		Birth:  "2022.2.22",
		Salary: 8000.0,
		Skill:  "牛魔拳",
		shadow: true,
	}

	//序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化错误，err= %v\n", err)
	}
	fmt.Printf("monster 序列化后结果=%v", string(data))

}

func testMap() {
	var a map[string]interface{} = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "洪崖洞"

	//序列化map
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误，err= %v\n", err)
	}
	fmt.Printf("a 序列化后结果=%v", string(data))
}

/*func testBaseType(){
//对基本数据类型序列化意义不大
}*/
func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{} = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = 7
	m1["address"] = "北京"

	var m2 map[string]interface{} = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = 8
	m2["address"] = "纽约"
	slice = append(slice, m1, m2)
	//序列化slice
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误，err= %v\n", err)
	}
	fmt.Printf("slice 序列化后结果=%v", string(data))
}

func main() {
	testStruct()
	testMap()
	testSlice()
}

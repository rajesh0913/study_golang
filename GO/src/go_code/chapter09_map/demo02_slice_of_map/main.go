package main

import "fmt"

//monster:记录monster信息（name、age）
func main() {
	monsters := make([]map[string]string, 2)
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "玉兔精"
		monsters[1]["age"] = "400"
	}
	newMonster := map[string]string{
		"name": "新妖怪", "age": "200",
	}
	monsters = append(monsters, newMonster)
	fmt.Println(monsters)
}

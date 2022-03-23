package main

import "fmt"

func modifyUser(user map[string]map[string]string, name string) {
	// for index,value:=range user{
	// 	if index==name{
	// 		user[index]["pwd"]="888888"
	// 		break
	// 	}else if
	// }
	if user[name] != nil {
		user[name]["pwd"] = "888888"
	} else {
		user[name] = make(map[string]string) //重要
		user[name]["nickname"] = "nickname"
		user[name]["pwd"] = "888888"
	}
}
func main() {
	//var user map[string]map[string]string
	user := make(map[string]map[string]string)
	modifyUser(user, "qinzhibao")
	fmt.Println(user)
	user["fengyifan"] = make(map[string]string)
	user["fengyifan"]["nickname"] = "fanfan"
	user["fengyifan"]["pwd"] = "777777"
	modifyUser(user, "fengyifan")
	fmt.Println(user)

}

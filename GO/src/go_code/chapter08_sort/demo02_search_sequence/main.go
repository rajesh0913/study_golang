package main

import "fmt"

func main() {
	names := [4]string{"白眉", "紫衫", "青翼", "金毛"}
	var hero string
	fmt.Println("请输入查找角色：")
	fmt.Scanln(&hero)
	//顺序查找1
	for i := 0; i < len(names); i++ {
		if hero == names[i] {
			fmt.Printf("找到%v,下标为%v\n", hero, i)
			break
		} else if i == len(names)-1 {
			fmt.Printf("名单中没有%v\n", hero)
		}
	}
	//顺序查找2
	index := -1
	for i := 0; i < len(names); i++ {
		if hero == names[i] {
			index = i
			//fmt.Printf("找到%v,下标为%v\n", hero, i)
			break
		}
	}
	if index != -1 {
		fmt.Printf("找到%v,下标为%v~\n", hero, index)
	} else {
		fmt.Printf("名单中没有%v~\n", hero)
	}
}

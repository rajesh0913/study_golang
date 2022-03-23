package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	map1 := make(map[int]int, 10)
	//按照key的字典序从小到大排列
	for i := 0; i < 10; i++ {
		map1[i] = rand.Intn(100) + 1
	}
	fmt.Println(map1)
	// 旧版，不升序，可以：
	// 1.将key放入切片
	// 2.对切片进行排序
	// 3.遍历切片输出

}

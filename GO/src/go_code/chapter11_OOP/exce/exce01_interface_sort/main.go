package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//实现对hero结构体切片排序
type Hero struct {
	Name string
	Age  int
}
type HeroSlice []Hero

//实现interface
func (t HeroSlice) Len() int {
	return len(t)
}
func (t HeroSlice) Less(i int, j int) bool {
	return t[i].Age < t[j].Age
}
func (t HeroSlice) Swap(i int, j int) {
	temp := t[i]
	t[i] = t[j]
	t[j] = temp
	//this[i],this[j]=this[j],this[i]
}

func main() {
	rand.Seed(time.Now().Unix())
	//对int类型切片排序
	var intSlice = []int{0, -99, 10, 5, 99}
	sort.Ints(intSlice)
	fmt.Println("sort.Ints:\n", intSlice)

	//对结构体切片排序
	var Heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("hero-%v", rand.Intn(100)),
			Age:  rand.Intn(500),
		}
		Heroes = append(Heroes, hero)
	}

	for _, value := range Heroes {
		fmt.Printf("%v\t", value)
	}
	fmt.Println()
	//sort
	sort.Sort(Heroes)
	for _, value := range Heroes {
		fmt.Printf("%v\t", value)
	}
}

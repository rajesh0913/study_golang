package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	//通过go向redis读写数据
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.dial err = ", err)
		return
	}
	fmt.Println("conn succ...", conn)

	//预处理关闭
	defer conn.Close()
	//写入数据
	_, err = conn.Do("set", "name", "tomjerry")
	if err != nil {
		fmt.Println("set err = ", err)
		return
	}
	fmt.Println("set done")
	//给数据设置有效时间
	_, err = conn.Do("expire", "name", 10)
	if err != nil {
		fmt.Println("expire err = ", err)
		return
	}
	//读取数据
	r, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("get err = ", err)
		return
	}
	//返回的r为interface{}，需要转换

	fmt.Println("get done\n name = ", r)
}

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
	_, err = conn.Do("HSet", "user01", "name", "tomjerry")
	if err != nil {
		fmt.Println("hash set err = ", err)
		return
	}

	_, err = conn.Do("HSet", "user01", "age", "18")
	if err != nil {
		fmt.Println("hash set err = ", err)
		return
	}

	/*
			批量操作

		_, err = conn.Do("HMSet", "user01", "name","tom","age", "18")
		if err != nil {
			fmt.Println("hash mset err = ", err)
			return
		}

		r1, err := redis.String(conn.Do("HMGet", "user01", "name","age"))
		if err != nil {
			fmt.Println("hash mget err = ", err)
			return
		}
		for i, v :=range r1{
			fmt.Printf("r[%d] = %s",i,v)
		}

	*/

	fmt.Println("hash set done")
	//读取数据
	r1, err := redis.String(conn.Do("HGet", "user01", "name"))
	if err != nil {
		fmt.Println("hash get err = ", err)
		return
	}
	r2, err := redis.Int(conn.Do("HGet", "user01", "age"))
	if err != nil {
		fmt.Println("hash get err = ", err)
		return
	}
	//返回的r为interface{}，需要转换

	fmt.Printf("get done\n name = %v, age = %v", r1, r2)
}

package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

/*链接池
1.事先初始化一定数量的链接放入链接池
2.go需要操作redis，直接从链接池取出链接使用
3.节省时间提高效率
*/
//定义全局pool
var pool *redis.Pool

//启动程序即初始化
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   //最大空闲链接数
		MaxActive:   0,   //表示和数据库的最大链接数，0为无限制
		IdleTimeout: 100, //最大空闲时间
		Dial: func() (redis.Conn, error) { //初始化链接的代码
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}
func main() {
	//链接池取出链接
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("set", "name", "tom")
	if err != nil {
		fmt.Println("conn.Do err = ", err)
		return
	}

	//取出
	r, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("conn.Do get err = ", err)
		return
	}
	fmt.Println("get done\nname = ", r)

	//关闭链接池
	pool.Close()
}

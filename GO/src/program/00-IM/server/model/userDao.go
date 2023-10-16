package model

import (
	"GO/src/program/00-IM/pkg/common/message"
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// 服务器启动后，使用UserDao实例全局变量访问redis
var (
	MyUserDao *UserDao
)

// 使用结构体方法进行Data access operation操作
type UserDao struct {
	pool *redis.Pool
}

// 使用工厂模式创建UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

// 1.根据id获取用户实例和错误
func (u *UserDao) getUserById(conn redis.Conn, id int) (user *message.User, err error) {
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		// users hash表中没有这个id
		if err == redis.ErrNil {
			err = ErrorUserNotExists
		}
		return
	}
	// 必须先创建user实例才能正确接收
	// user = &User{}
	err = json.Unmarshal([]byte(res), &user)

	if err != nil {
		fmt.Println("json unmarshal fail error:", err)
		return
	}
	return
}

// 2.登录校验
func (u *UserDao) Login(userId int, userPwd string) (user *message.User, err error) {
	// 连接池
	conn := u.pool.Get()
	defer conn.Close()
	user, err = u.getUserById(conn, userId)
	if err != nil {
		return
	}
	// 获取到信息
	if userPwd != user.UserPwd {
		err = ErrorUserPwd
		return
	}
	return
}

// 3. 注册
func (u *UserDao) Register(user *message.User) (err error) {
	// 连接池
	conn := u.pool.Get()
	defer conn.Close()
	// 检验用户id是否存在
	_, err = u.getUserById(conn, (*user).UserId)
	if err != ErrorUserNotExists {
		fmt.Println("(*user).UserId:", (*user).UserId)
		fmt.Println("error: ", err)
		err = ErrorUserExists
		return
	}
	// 用户id不存在，开始注册
	data, err := json.Marshal((*user))
	if err != nil {
		return
	}
	_, err = conn.Do("HSet", "users", (*user).UserId, string(data))
	if err != nil {
		fmt.Println("register conn Register fail error: ", err)
		return
	}
	return
}

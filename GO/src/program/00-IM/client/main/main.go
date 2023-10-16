/*
 */
package main

import (
	"GO/src/program/00-IM/client/process"
	"fmt"
	"os"
)

/*
UserId: user id
UserPwd:user password
*/
var userId int
var userPwd, userName string

func main() {
	/*
		key:接收用户选择
		loop:判断是否继续选择
	*/
	var key int
	var loop bool = true
	for loop {
		fmt.Println("----------欢迎登录多人聊天系统----------")
		fmt.Println("\t\t 1 登录聊天室")
		fmt.Println("\t\t 2 注册用户")
		fmt.Println("\t\t 3 退出聊天室")
		fmt.Println("\t\t 请选择(1-3):")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户ID:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码:")
			fmt.Scanf("%s\n", &userPwd)

			// 登录操作
			// 1.创建up
			up := &process.UserProcess{}
			// 2.调用登录函数
			_ = up.Login(userId, userPwd)

		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码:")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户名:")
			fmt.Scanf("%s\n", &userName)
			// 注册
			// 1. 创建用户登录实例
			up := &process.UserProcess{}
			// 2. 调用注册函数
			_ = up.Register(userId, userPwd, userName)

		case 3:
			fmt.Println("退出聊天室")
			os.Exit(0)
		default:
			fmt.Println("输入有误，请重新输入")
			loop = true
		}

	}

}

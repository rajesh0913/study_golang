package process

import (
	"GO/src/program/00-IM/client/model"
	"GO/src/program/00-IM/pkg/common/message"
	"fmt"
)

// 登录成功之后初始化
var (
	curUser *model.CurUser
)

// 客户端维护的在线用户map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 1024)

// 客户端显示当前在线用户
func outputOnlineUser() {
	fmt.Println("当前在线用户列表:")
	for id := range onlineUsers {
		fmt.Println("用户id:\t", id)
	}
}

// 处理返回的notifyonlineusermes消息
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.UserStatus
	onlineUsers[notifyUserStatusMes.UserId] = user

	outputOnlineUser()

}

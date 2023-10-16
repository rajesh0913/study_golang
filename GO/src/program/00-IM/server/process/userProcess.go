package process

import (
	"GO/src/program/00-IM/pkg/common/message"
	"GO/src/program/00-IM/server/model"
	"GO/src/program/00-IM/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int
}

// 处理用户登录请求的函数
func (u *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	// 1.先从mes中取出data，反序列化为LoginMes结构
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte((*mes).Data), &loginMes)
	if err != nil {
		fmt.Printf("json.Unmarshal fail err: %v", err)
		return
	}
	// 3.声明resMes\loginResMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	var loginResMes message.LoginResMes

	// 2.验证用户是否合法
	// 使用myUserDao实例到redis进行验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ErrorUserNotExists {

			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ErrorUserPwd {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误"
		}
	} else {
		loginResMes.Code = 200
		fmt.Println(user.UserId, "登录成功...")
		// 登录成功，放入在线用户列表
		(*u).UserId = loginMes.UserId
		(*userMgr).AddOnlineUser(u)
		u.NotifyOtherOnlineUser((*u).UserId)
		// 返回消息加入在线用户消息
		for id := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}
	}
	// 4.准备返回的信息序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Printf("json.Marshal fail err: %v", err)
		return
	}
	resMes.Data = string(data)
	// 5.将resMes进行序列化并发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Printf("json.Marshal fail err: %v", err)
		return
	}
	// 6.发送data,需要一个Transfer结构体
	tf := &utils.Transfer{
		Conn: u.Conn,
	}
	err = tf.WritePkg(data)
	return
}

// 处理用户注册请求的函数
func (u *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	// 1. 取出data
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("ServerProcessRegister json.Unmashal mes.Data fail error: ", err)
		return
	}
	// 2. 创建返回信息
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes

	// 3. redis服务器注册
	err = model.MyUserDao.Register((*message.User)(&registerMes.User))
	if err != nil {
		if err == model.ErrorUserExists {
			registerResMes.Code = 505
			registerResMes.Error = err.Error()
		} else {
			registerResMes.Code = 404
			registerResMes.Error = "ErrorRegiserUnkownError"
		}
	} else {
		registerResMes.Code = 200
	}
	// 4. 准备返回信息
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Printf("json.Marshal registerResMes fail err: %v", err)
		return
	}
	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Printf("json.Marshal register resMes fail err: %v", err)
		return
	}
	// 5.发送data,需要一个Transfer结构体
	tf := &utils.Transfer{
		Conn: u.Conn,
	}
	err = tf.WritePkg(data)

	return
}

// 通知用户上线
func (u *UserProcess) NotifyOtherOnlineUser(userId int) {
	for id, up := range (*userMgr).onlineUsers {
		if id == (*u).UserId {
			continue
		}
		up.NoticeMeOnline(userId)
	}
}

// notice
func (u *UserProcess) NoticeMeOnline(userId int) {
	// 准备notice消息
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.UserStatus = message.UserOnline

	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json marshal notifyUserStatusMes fail error: ", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json marshal notifyUser mes fail error: ", err)
		return
	}
	// 发送消息
	tf := &utils.Transfer{
		Conn: (*u).Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("notifyMeOnline tf.WritePkg fail error: ", err)
		return
	}

}

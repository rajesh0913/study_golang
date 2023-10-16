package process

import (
	"GO/src/program/00-IM/client/model"
	"GO/src/program/00-IM/client/utils"
	"GO/src/program/00-IM/pkg/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type UserProcess struct {
}

func (u *UserProcess) Login(userId int, userPwd string) (err error) {

	// fmt.Printf("userId = %v,userPwd = %s\n", userId, userPwd)
	// return nil
	// 1.连接到服务器
	conn, err := net.Dial("tcp", "192.168.5.9:8889")
	if err != nil {
		fmt.Printf("net.dial err: %v.", err)
		return
	}
	defer conn.Close()
	// 2.通过conn发送消息
	var mes message.Message
	mes.Type = message.LoginMesType

	// 3.创建loginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	// 4.将loginMes结构体序列化

	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Printf("json.marshal err: %v.", err)
		return
	}
	// 5.将序列化后的消息传递给mes
	mes.Data = string(data)

	// 6.将mes结构体序列化,得到要发送的消息data
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Printf("json.marshal err: %v.", err)
		return
	}
	/*
		7.1 将消息长度发送给服务器
		net.write发送的数据类型是byte切片
		需要现将消息的长度转换为一个表示长度的byte切片
	*/
	var pkgLen uint32 = uint32(len(data))
	// 32位，4个字节长度
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Printf("conn.write(buf) err: %v.", err)
		return
	}
	// fmt.Printf("客户端发送消息长度:%v,内容:%v\n", len(data), string(data))
	// 7.2 发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Printf("conn.write(data) err: %v.", err)
		return
	}
	// 处理返回的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("message.readPkg fail error: ", err)
		return
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		// 登录成功
		// 初始化curUser
		curUser = &model.CurUser{}
		curUser.Conn = conn
		curUser.UserId = userId
		curUser.UserStatus = message.UserOnline
		// 显示在线用户列表
		fmt.Println("当前在线用户：")
		for _, v := range loginResMes.UsersId {
			if v == userId {
				continue
			}
			fmt.Println("用户id:\t", v)

			// 完成onlineUsers初始化
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user

		}
		fmt.Printf("\n\n")
		// 1. 启动协程，保持两端同步
		go serverProcessMes(conn)
		// 2. 显示登录菜单
		for {
			showMenu()
		}

	} else {
		fmt.Println(loginResMes.Error)
	}
	return
}

func (u *UserProcess) Register(userId int, userPwd string,
	userName string) (err error) {
	// 1. 连接服务器
	conn, err := net.Dial("tcp", "192.168.5.9:8889")
	if err != nil {
		fmt.Printf("net.dial err: %v.", err)
		return
	}
	defer conn.Close()

	// 2. 准备发送信息
	var mes message.Message
	mes.Type = message.RegisterMesType

	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	// 3. 序列化注册信息

	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal registerMes fail error:", err)
		return
	}

	mes.Data = string(data)
	// 4. 序列化发送信息
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal mes fail error:", err)
		return
	}

	// 5. 发送信息长度,使用utils的Writepkg函数
	tf := utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("RegisterMes tf.WritePkg mes fail error:", err)
		return
	}

	// 6. 处理返回消息
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("RegisterMes tf.ReadPkg mes fail error:", err)
		return
	}
	// 处理信息，创建返回注册信息实例
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，请重新登录！")
	} else {
		fmt.Println(registerResMes.Error)
	}
	os.Exit(0)
	return
}

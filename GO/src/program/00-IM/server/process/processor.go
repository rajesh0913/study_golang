package process

import (
	"GO/src/program/00-IM/pkg/common/message"
	"GO/src/program/00-IM/server/utils"
	"fmt"
	"io"
	"net"
)

// 创建结构体

type Processor struct {
	Conn net.Conn
}

/*
编写一个serverProcessMes函数
根据客户端发送消息类型，决定调用不同函数处理
*/
func (s *Processor) ServerProcessMes(mes *message.Message) (err error) {
	switch (*mes).Type {
	// 处理登录
	case message.LoginMesType:
		// 处理登录问题
		// 创建一个UserProcess实例
		up := &UserProcess{
			Conn: (*s).Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		// 处理登录问题
		// 创建一个UserProcess实例
		up := &UserProcess{
			Conn: (*s).Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		smsProcess := &SmsProcess{}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在，无法处理...")
		fmt.Println("mes = ", mes)

	}
	return
}

func (p *Processor) Process2() (err error) {
	tf := &utils.Transfer{
		Conn: (*p).Conn,
	}
	for {
		// 将数据包读取直接封装成一个函数，返回message和err
		message, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Printf("客户端退出通讯,服务器也退出\n")
				return err
			}
			fmt.Printf("readPkg() err:%v\n", err)
			return err
		}
		err = (*p).ServerProcessMes((&message))
		if err != nil {
			fmt.Println("serverProcessMes fail error: ", err)
			return err
		}
	}
}

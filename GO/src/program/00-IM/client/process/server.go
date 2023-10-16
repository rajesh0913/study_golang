package process

import (
	"GO/src/program/00-IM/client/utils"
	"GO/src/program/00-IM/pkg/common/message"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func showMenu() {
	fmt.Println("----- 恭喜xxx登录成功 -----")
	fmt.Println("----- 1. 显示在线用户列表 -----")
	fmt.Println("----- 2. 发送信息 -----")
	fmt.Println("----- 3. 信息列表 -----")
	fmt.Println("----- 4. 退出系统 -----")
	fmt.Println("请选择(1-4):")
	var key int
	var content string
	fmt.Scanf("%d\n", &key)

	// 定义smsprocess实例
	smsProcess := &SmsProcess{}
	switch key {
	case 1:
		outputOnlineUser()
	case 2:
		fmt.Println("输入消息:")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("输入错误，请重新输入...")
	}
}
func serverProcessMes(conn net.Conn) {
	// 创建接发实例
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {

		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("serverProcessMes tf.readPkg fail err: ", err)
			return
		}
		// fmt.Printf("mes = %v", mes)
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			// 1. 处理消息
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			// 2. 将上线用户存入本地文件
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType:
			outputGroupMes(&mes)
		default:
			fmt.Println("服务器端返回未知类型消息...")
			fmt.Printf("mes = %v", mes)
		}
	}

}

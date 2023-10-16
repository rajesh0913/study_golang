package process

import (
	"GO/src/program/00-IM/client/utils"
	"GO/src/program/00-IM/pkg/common/message"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {
}

func (sp *SmsProcess) SendGroupMes(content string) (err error) {
	// 准备消息
	var mes message.Message
	mes.Type = message.SmsMesType
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = curUser.UserId
	smsMes.UserStatus = curUser.UserStatus
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json marshal smsMes fail error: ", err)
		return
	}

	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json marshal mes fail error: ", err)
		return
	}

	// transfer
	tf := &utils.Transfer{
		Conn: curUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes tf.WritePkg fail error: ", err)
		return
	}
	return
}

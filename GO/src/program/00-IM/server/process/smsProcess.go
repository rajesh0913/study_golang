package process

import (
	"GO/src/program/00-IM/client/utils"
	"GO/src/program/00-IM/pkg/common/message"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {
}

func (sp *SmsProcess) SendGroupMes(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("SendGoupMes json.Unmarshal mes.Data fail error", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGoupMes json.marshal mes fail error", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		if id == smsMes.UserId {
			continue
		}
		(*sp).SendMesToEach(data, up.Conn)
	}

}
func (sp *SmsProcess) SendMesToEach(data []byte, conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendMesToEach tf.WritePkg fail error: ", err)
	}
}

package process

import (
	"GO/src/program/00-IM/pkg/common/message"
	"encoding/json"
	"fmt"
)

func outputGroupMes(mes *message.Message) {
	// 1. 处理消息
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("outputGroupMes json.Unmarshal fail error: ", err)
		return
	}

	info := fmt.Sprintf("用户id:%d\t群发:\t%s", smsMes.UserId, smsMes.Content)
	fmt.Println(info)
	fmt.Println()

}

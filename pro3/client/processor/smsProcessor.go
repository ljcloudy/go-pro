package processor

import (
	"encoding/json"
	"fmt"
	"pro3/client/utils"
	"pro3/common/message"
)

type SmsProcessor struct {
}

//发送群聊的消息

func (smsProcessor *SmsProcessor) SendGroupMes(content string) (err error) {
	// 1. 创建mes
	var mes message.Message
	mes.Type = message.SmsMesType

	// 2. 创建SmsMes
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = curUser.UserId
	smsMes.UserStatus = curUser.UserStatus

	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//发送服务器
	tf := &utils.Transfer{
		Conn: curUser.Conn,
	}
	err = tf.WritePkg(data)
	return
}

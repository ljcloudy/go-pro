package processor

import (
	"encoding/json"
	"fmt"
	"net"
	"pro3/common/message"
	"pro3/server/utils"
)

type SmsProcess struct {
}

func (smsProcess *SmsProcess) SendGroupMes(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	data, err := json.Marshal(mes)

	for id, up := range userMgr.onlineUsers {
		if id == smsMes.UserId {
			continue
		}
		smsProcess.SendMesToEachOnlineUser(data, up.Conn)
	}
}
func (smsProcess *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败 err=", err)
	}
}

package processor

import (
	"encoding/json"
	"fmt"
	"pro3/common/message"
)

func outputGroupMes(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	//显示信息
	info := fmt.Sprintf("用户Id:%d\t-->\t%s", smsMes.UserId, smsMes.Content)
	fmt.Println(info)
	fmt.Println()
}

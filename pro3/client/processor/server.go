package processor

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"pro3/client/utils"
	"pro3/common/message"
)

func showMenu() {
	fmt.Println("-------恭喜XXX登录成功---------")
	fmt.Println("-------1 显示在线用户列表---------")
	fmt.Println("-------2 发送消息---------")
	fmt.Println("-------3 信息列表---------")
	fmt.Println("-------4 退出系统---------")
	fmt.Println("请选择(1-4):")

	var key int
	var content string
	fmt.Scanf("%d\n", &key)

	smsProcessor := &SmsProcessor{}
	switch key {
	case 1:
		//fmt.Println("显示在线用户")
		outputOnlineUser()
	case 2:
		fmt.Println("你想对大家说：")
		fmt.Scanf("%s\n", &content)
		smsProcessor.SendGroupMes(content)
	case 3:
		fmt.Println("消息列表")
	case 4:
		fmt.Println("你选择退出系统!")
		os.Exit(0)
	default:
		fmt.Println("您的输入不正确！")
	}
}

//和服务器保持通讯
func serverProcessMes(conn net.Conn) {
	//创建transfer实例，不停的读取服务器发送的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务器的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err=", err)
			return
		}
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			//fmt.Println("通知用户上线消息=", notifyUserStatusMes)
			updateUserStatus(notifyUserStatusMes)
		case message.SmsMesType:
			outputGroupMes(&mes)
		default:
			fmt.Println("服务器返回了未知消息！")
		}
		fmt.Printf("mes=%v", mes)
	}
}

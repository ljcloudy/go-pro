package main

import (
	"fmt"
	"io"
	"net"
	"pro3/common/message"
	"pro3/server/processor"
	"pro3/server/utils"
)

type Process struct {
	Conn net.Conn
}

func (process *Process) processMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登录逻辑
		up := &processor.UserProcessor{
			Conn: process.Conn,
		}
		err = up.ProcessLogin(mes)
	case message.RegisterMesType:
		//处理注册
		up := &processor.UserProcessor{
			Conn: process.Conn,
		}
		err = up.ProcessRegister(mes)
	case message.SmsMesType:
		smsProcess := &processor.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在，无法处理。。。")
	}
	return
}

func (process *Process) process() (err error) {
	for {
		tf := &utils.Transfer{
			Conn: process.Conn,
		}
		mes, err := tf.ReadPkg()
		fmt.Println("mes=", mes)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务端也退出.")
				return err
			} else {
				fmt.Println("readPkg err=", err)
				return err
			}
		}
		err = process.processMes(&mes)
		if err != nil {
			return err
		}
	}

}

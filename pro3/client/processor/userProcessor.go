package processor

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"pro3/client/utils"
	"pro3/common/message"
)

type UserProcessor struct {
}

func (this *UserProcessor) Login(userId int, userPwd string) (err error) {
	//1. 连接服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}

	defer conn.Close()

	//2. 准备给服务发送消息
	var mes message.Message
	mes.Type = message.LoginMesType
	//3. 创建一个LoginMesType
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//4. 将loginMes 序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	// 5. 将data赋值
	mes.Data = string(data)
	//6. 将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	// 7 发送data
	tf := &utils.Transfer{
		Conn: conn,
	}
	fmt.Println("send data=", string(data))
	tf.WritePkg(data)

	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg  err=", err)
		return
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		curUser.Conn = conn
		curUser.UserId = userId
		curUser.UserStatus = message.UserOnline
		fmt.Println("当前在线用户列表如下：")
		for _, v := range loginResMes.UserId {
			if v == userId {
				continue
			}
			fmt.Println("用户Id:\t", v)
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Print("\n\n")
		//该协程保持和服务器的通讯。如果服务器有数据推送客户端
		go serverProcessMes(conn)
		for {
			showMenu()
		}
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	} else {
		fmt.Println(loginResMes.Error)
	}

	return
}

func (this *UserProcessor) Register(userId int, userPwd string, userName string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer conn.Close()

	var mes message.Message
	mes.Type = message.RegisterMesType

	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	data, err := json.Marshal(registerMes)
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
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息失败err=", err)
	}

	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg err=", err)
		return
	}
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)

	if registerResMes.Code == 200 {
		fmt.Println("注册成功！，请重新登录！")
		os.Exit(0)
	} else {
		fmt.Println("注册失败，" + registerResMes.Error)
		os.Exit(0)
	}
	return

}

//{\"userId\":100,\"userPwd\":\"123456\",\"userName\":\"lijianyun\"}

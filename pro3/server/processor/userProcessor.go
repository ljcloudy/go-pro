package processor

import (
	"encoding/json"
	"fmt"
	"net"
	"pro3/common/message"
	"pro3/server/model"
	"pro3/server/utils"
)

type UserProcessor struct {
	Conn   net.Conn
	UserId int
}

func (up *UserProcessor) ProcessLogin(mes *message.Message) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	//返回响应
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	var loginResMes message.LoginResMes

	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误！"
		}
	} else {
		loginResMes.Code = 200
		fmt.Println("登录成功user=", user)
		up.UserId = loginMes.UserId
		userMgr.AddOnlineUser(up)
		//通知其他用户，新用户上线
		up.NotifyOthersOnlineUser(loginMes.UserId)
		//将当前在线用户的ID，放入LoginResMes.UserId
		for id := range userMgr.onlineUsers {
			loginResMes.UserId = append(loginResMes.UserId, id)
		}
	}

	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail, err=", err)
		return
	}
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)

	tf := &utils.Transfer{
		Conn: up.Conn,
	}
	tf.WritePkg(data)
	return
}

func (up *UserProcessor) ProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes

	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	//响应
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes

	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册未知错误！"
		}
	} else {
		registerResMes.Code = 200
	}

	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	tf := &utils.Transfer{
		Conn: up.Conn,
	}
	err = tf.WritePkg(data)
	return
}

func (up *UserProcessor) NotifyOthersOnlineUser(userId int) {
	for id, up := range userMgr.onlineUsers {
		if id == userId {
			continue
		}
		up.NotifyMeOnline(userId)
	}
}
func (up *UserProcessor) NotifyMeOnline(userId int) {
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	data, err := json.Marshal(notifyUserStatusMes)
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
		Conn: up.Conn,
	}
	fmt.Println("通知其他用户上线，", string(data))
	tf.WritePkg(data)

}

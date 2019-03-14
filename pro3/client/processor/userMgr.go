package processor

import (
	"fmt"
	"pro3/client/model"
	"pro3/common/message"
)

//客户端要维护map
var onlineUsers = make(map[int]*message.User, 10)
var curUser model.CurUser

func outputOnlineUser() {
	fmt.Println("当前在线用户列表：")
	for id := range onlineUsers {
		fmt.Println("用户Id:\t", id)
	}
}
func updateUserStatus(mes message.NotifyUserStatusMes) {
	user, ok := onlineUsers[mes.UserId]
	if !ok {
		user = &message.User{
			UserId: mes.UserId,
		}
	}
	user.UserStatus = mes.Status
	onlineUsers[mes.UserId] = user

	outputOnlineUser()
}

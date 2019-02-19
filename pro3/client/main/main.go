package main

import (
	"fmt"
	"os"
	"pro3/client/processor"
)

var userId int
var userPwd string

func main() {
	//接受用户选择
	var key int
	//判断是否还继续显示菜单

	var loop = true
	up := &processor.UserProcessor{}
	for loop {
		fmt.Println("--------------欢迎登录多人聊天系统-----------------")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3):")

		fmt.Scanf("%d\n", &key)

		switch key {

		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户ID")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入密码")
			fmt.Scanf("%s\n", &userPwd)

			up.Login(userId, userPwd)

			//loop = false
		case 2:
			var userName, userPwd string
			var userId int

			fmt.Printf("请输入用户Id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码:")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户名:")
			fmt.Scanf("%s\n", &userName)

			up.Register(userId, userPwd, userName)
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("您的输入有误，请重新输入！")
		}
	}

}

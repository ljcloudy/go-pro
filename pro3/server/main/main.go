package main

import (
	"fmt"
	"net"
	"pro3/server/model"
	"time"
)


func initUserDao(){
	model.MyUserDao = model.NewUserDao(pool)
}
func main() {

	initPool("localhost:6379", 16,0,300*time.Second)
	initUserDao()


	fmt.Println("服务器在8889端口监听....")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}

	//一旦监听成功，就等待客户端来连接服务器
	for {
		fmt.Println("等待客户端连接....")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}
		//一旦连接成功， 则启动一个协程和客户端保持通讯
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	process := &Process{
		Conn: conn,
	}
	err := process.process()
	if err != nil {
		fmt.Println("客户端和服务器通讯协程错误err=", err)
		return
	}

}

package model

import (
	"net"
	"pro3/common/message"
)

type CurUser struct {
	Conn net.Conn
	message.User
}

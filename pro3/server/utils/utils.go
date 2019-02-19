package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"pro3/common/message"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	buf := make([]byte, 8096)

	fmt.Println("读取客户端发送的数据....")
	_, err = this.Conn.Read(buf[:4])
	if err != nil {
		fmt.Println("conn Read err=", err)
		return
	}
	//根据buf[:4] 转成一个uint32
	var pkgLen uint32 = binary.BigEndian.Uint32(buf[:4])

	//根据pkgLen 读取消息内容
	n, err := this.Conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		return
	}
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json Unmarshal err=", err)
		return
	}
	return
}
func (this *Transfer) WritePkg(data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	n, err := this.Conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write fail ", err)
		return
	}
	_, err = this.Conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write  err=", err)
		return
	}
	return
}

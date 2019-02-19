package message

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMesType"
	RegisterResMesType = "RegisterResMesType"
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息体
}

type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMes struct {
	Code  int    `json:"code"` //返回状态码 ：500 表示该用户未注册 200登录成功
	Error string `json:"error"`
}

// 用户结构体
type User struct {
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}

type RegisterMes struct {
	User User `json:"user"` //类型就是User结构体
}

type RegisterResMes struct{
	Code int  `json:"code"`
	Error string `json:"error"`
}

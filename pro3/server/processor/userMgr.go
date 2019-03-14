package processor

import "fmt"

var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcessor
}

//初始化UserMgr
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcessor, 1024),
	}
}

func (userMgr *UserMgr) AddOnlineUser(up *UserProcessor) {
	userMgr.onlineUsers[up.UserId] = up
}

func (userMgr *UserMgr) DeleteOnlineUser(up *UserProcessor) {
	delete(userMgr.onlineUsers, up.UserId)
}
func (userMgr *UserMgr) GetAllOnlineUser() map[int]*UserProcessor {
	return userMgr.onlineUsers
}
func (userMgr *UserMgr) GetOnlineUserById(userId int) (up *UserProcessor, err error) {
	up, ok := userMgr.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("用户%d不存在", userId)
		return
	}
	return
}

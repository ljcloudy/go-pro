package model

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"pro3/common/message"
)

var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

func NewUserDao(pool *redis.Pool) *UserDao {
	userDao := &UserDao{
		pool: pool,
	}
	return userDao
}

func (userDao *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
	}
	return
}

// login
func (userDao *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	conn := userDao.pool.Get()
	defer conn.Close()

	user, err = userDao.getUserById(conn, userId)
	if err != nil {
		return
	}
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

func (userDao *UserDao) Register(user *message.User) (err error) {
	conn := userDao.pool.Get()
	defer conn.Close()

	_, err = userDao.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}
	data ,err := json.Marshal(user)
	if err != nil{
		return
	}

	_,err = conn.Do("HSet", "users",user.UserId, string(data))
	if err != nil{
		fmt.Println("保存注册用户err=", err)
		return
	}
	return
}

package model

import (
	"net"
	"chatroom/common/message"
)
//在客户端很多地方会使用到curUser,将其作为一个全局
type CurUser struct {
	Conn net.Conn
	message.User
} 
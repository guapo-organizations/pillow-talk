package client

import (
	"github.com/gorilla/websocket"
	"github.com/guapo-organizations/pillow-talk/hub"
)

//连接的客户定义

type Client struct {
	//客户在哪个房间
	hub *hub.Hub

	//连接客户的句柄，也就是给这个用户或者读取这个用户信息的输入输出流
	conn *websocket.Conn

	//房间中的聊天内容
	send chan []byte
}

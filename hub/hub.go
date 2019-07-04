package hub

import "github.com/guapo-organizations/pillow-talk/client"

//聊天的房间定义
//聊天室的作用，是把一个用户发送的信息通知给其他在这个房间的用户

type Hub struct {
	//房间中的客户
	clients map[*client.Client]bool

	//通知给用户新信息
	broadcast chan []byte

	//新加入房间的用户
	register chan *client.Client

	//从房间离开的用户
	unregister chan *client.Client
}

//创建一个房间
func NewHub() *Hub {
	return &Hub{
		clients: make(map[*client.Client]bool),
		//为什么定义为无缓冲chan,作为广播呢?，因为每次只需通知一条信息
		broadcast: make(chan []byte),
		//为什么定义为无缓冲chan，作为注册呢？,因为每次连接只有一个用户进来连接,一个连接对应一个用户
		register: make(chan *client.Client),
		//为什么定义为无缓冲chan呢？,因为每次离开房间的用户只有一个，一个连接对应一个用户
		unregister: make(chan *client.Client),
	}
}

func (h *Hub) run() {

}

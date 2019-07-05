package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	//聊天室端口
	addr = flag.String("addr", ":8080", "http service address")
)

//前端websocket逻辑
func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func main() {
	flag.Parse()
	//创建一个房间
	hub := newHub()
	//执行房间的，信息转发，用户加入离开逻辑
	go hub.run()

	//前端websocket，这个可以找一个专业的前端人员配合重写
	http.HandleFunc("/", serveHome)

	//后端ws逻辑，新用户加入
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		//新用户加入
		serveWs(hub, w, r)
	})

	//启动服务咯
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("启动服务失败:", err)
	}
}

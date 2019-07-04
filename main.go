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

func main() {
	flag.Parse()
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("启动服务失败:", err)
	}
}

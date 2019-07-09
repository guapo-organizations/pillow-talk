# 聊天室

一个自己团队的聊天工具



## 安装
1. 执行 compile.bat
2. 将编译好的linux可执行文件和home.html放同一个目录
3. 运行pillow-talk addr :8080

```
pillow-talk addr :ws监听的端口
```

# 写代码时候领悟到websocket的一些知识

## websocket ping pong 机制

首先是百度一堆说的 ping pong 机制是为了保持用户长时间不动也不断开连接的机制

### 服务端代码怎么写？

#### pong
```
func reader(ws *websocket.Conn) {
	defer ws.Close()
	ws.SetReadLimit(512)
	ws.SetReadDeadline(time.Now().Add(pongWait))
	//为什么读取的是pong呢？  读取嘛，肯定是客户端发送的数据，ping也是客户端发送过来的，所以这里是pong
	ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			break
		}
	}
}
```
为什么要在读这里写pong？  因为读是客户端的写，

ping的话肯定是通过客户端的写传进来的，所以，服务端的读设置的是pong，来响应客户端的ping


#### ping
```
func writer(ws *websocket.Conn, lastMod time.Time) {
	lastError := ""
	//定时给那边发送ping
	pingTicker := time.NewTicker(pingPeriod)
	//定时检查文件
	fileTicker := time.NewTicker(filePeriod)
	defer func() {
		pingTicker.Stop()
		fileTicker.Stop()
		ws.Close()
	}()

	for {
		select {
		//定时读取文件,然后返回
		case <-fileTicker.C:
			var p []byte
			var err error

			p, lastMod, err = readFileIfModified(lastMod)

			if err != nil {
				if s := err.Error(); s != lastError {
					lastError = s
					p = []byte(lastError)
				}
			} else {
				lastError = ""
			}

			if p != nil {
				//写入文件
				ws.SetWriteDeadline(time.Now().Add(writeWait))
				if err := ws.WriteMessage(websocket.TextMessage, p); err != nil {
					return
				}
			}
		case <-pingTicker.C:
			//定时发送ping
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}
```

同理可证，服务端的写，设置的是ping
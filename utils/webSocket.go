package utils

import (
	"github.com/gorilla/websocket"
	"net/http"
)

func GetWsCoon(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*websocket.Conn, error) {
	Upgrader := websocket.Upgrader{
		// 读取缓存大小
		ReadBufferSize: 1024,
		// 写入缓存大小
		WriteBufferSize: 1024,
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	//升级http请求
	return Upgrader.Upgrade(w, r, responseHeader)
}

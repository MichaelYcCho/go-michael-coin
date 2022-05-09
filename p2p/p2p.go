package p2p

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/michael_cho77/go-michael-coin/utils"
)

var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	// Port :3000 will upgrade the request from :4000
	openPort := r.URL.Query().Get("openPort")
	ip := utils.Spliter(r.RemoteAddr, ":", 0)
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return openPort != "" && ip != ""
	}

	conn, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)
	initPeer(conn, ip, openPort)
	time.Sleep(20 * time.Second)
	conn.WriteMessage(websocket.TextMessage, []byte("Hello, from Port 3000"))
}

func AddPeer(address, port, openPort string) {
	// Port :4000 -> :3000 으로 Dial
	conn, _, err := websocket.DefaultDialer.Dial(
		fmt.Sprintf("ws://%s:%s/ws?openPort=%s", address, port, openPort[1:]), nil)
	utils.HandleErr(err)
	initPeer(conn, address, port)
	time.Sleep(10 * time.Second)
	conn.WriteMessage(websocket.TextMessage, []byte("Hello, from Port 4000"))
}

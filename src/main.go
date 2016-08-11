package main

import (
	"chat"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var room *chat.ChatRoom = chat.CreateChatRoom("The first")

var hub chat.Hub = chat.Hub{}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func upgradeWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	chatRoom := hub.Get(r.URL.Query().Get("room"))
	chatRoom.Register(conn, r.URL.Query().Get("name"))
}

func main() {
	m := chat.Message{"HEY YOU GUYS", "Bill Bobb", time.Now()}
	room.Broadcast(&m)
}

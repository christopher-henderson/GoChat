package main

import (
	"chat"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var hub chat.Hub = chat.CreateHub()

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func upgradeWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	chatRoom := hub.Get(r.URL.Query().Get("room"))
	chatRoom.Register(conn, r.URL.Query().Get("name"))
}

func main() {
	http.HandleFunc("/ws", upgradeWebsocket)
	http.Handle("/", http.FileServer(http.Dir("../static")))
	http.ListenAndServe(":8000", nil)
}

package main

import (
	"chat"
	"time"
)

// func newClient(ws *websocket.Conn) {
// 	request := ws.Request()
// 	chatRoom := request.URL.Query()
// 	room := getRoom(chatRoom)
// 	room.register(ws)
// }

var room *chat.ChatRoom = chat.CreateChatRoom()

func main() {
	m := chat.Message{"HEY YOU GUYS", "Bill Bobb", time.Now()}
	room.Broadcast(&m)
}

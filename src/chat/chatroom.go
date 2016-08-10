package chat

import (
	"fmt"

	"golang.org/x/net/websocket"
)

type ChatRoom struct {
	clients map[*Client]bool
	log     *Log
}

func CreateChatRoom() *ChatRoom {
	log := Log{}
	room := ChatRoom{make(map[*Client]bool), &log}
	return &room
}

func (room *ChatRoom) Register(clientSocket *websocket.Conn) {
	client := &Client{clientSocket, room}
	if _, ok := room.clients[client]; ok {
		return
	}
	room.clients[client] = true
}

func (room *ChatRoom) Unregister(client *Client) {
	if _, ok := room.clients[client]; !ok {
		return
	}
	delete(room.clients, client)
}

func (room *ChatRoom) Broadcast(message *Message) {
	room.log.Append(message)
	// for client := range room.clients {
	// 	client.broadcast(message)
	// }
	for _, m := range room.log.Slice() {
		fmt.Println("thing is ", m)
	}
}

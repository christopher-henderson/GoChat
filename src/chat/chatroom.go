package chat

import (
	"sync"

	"github.com/gorilla/websocket"
)

type ChatRoom struct {
	sync.Mutex
	clients map[string]*Client
	log     *Log
	Name    string
}

func CreateChatRoom(name string) *ChatRoom {
	log := Log{}
	room := ChatRoom{clients: make(map[string]*Client), log: &log, Name: name}
	return &room
}

func (room *ChatRoom) Register(clientSocket *websocket.Conn, name string) {
	room.Lock()
	defer room.Unlock()
	if _, ok := room.clients[name]; ok {
		return
	}
	room.clients[name] = CreateClient(clientSocket, room, name)
}

func (room *ChatRoom) Unregister(client string) {
	room.Lock()
	defer room.Unlock()
	if _, ok := room.clients[client]; !ok {
		return
	}
	delete(room.clients, client)
}

func (room *ChatRoom) Broadcast(message *Message) {
	room.log.Append(message)
	for _, client := range room.clients {
		go client.Accept(message)
	}
}

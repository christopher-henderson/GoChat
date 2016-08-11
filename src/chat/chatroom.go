package chat

import (
	"encoding/json"
	"fmt"
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
	fmt.Println("Registering " + name + " to chatroom " + room.Name)
	room.clients[name] = CreateClient(clientSocket, room, name)
	room.Broadcast(CreateMessage(name+" has joined the room.", "server"))
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
	m, err := json.Marshal(PrepareJSON(message))
	if err != nil {
		// NUTS!
		fmt.Println("Failed to marshal message.")
		return
	}
	room.BroadcastJSON(string(m))
}

func (room *ChatRoom) BroadcastJSON(message string) {
	for _, client := range room.clients {
		go client.Notify(message)
	}
}

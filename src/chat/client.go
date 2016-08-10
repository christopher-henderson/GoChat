package chat

import "golang.org/x/net/websocket"

type Client struct {
	ws       *websocket.Conn
	chatroom *ChatRoom
}

func CreateClient(ws *websocket.Conn, chatroom *ChatRoom) *Client {
	client := &Client{ws, chatroom}
	return client
}

func (c *Client) broadcast(message *Message) {

}

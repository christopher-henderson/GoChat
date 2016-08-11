package chat

import (
	"encoding/json"
	"io"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	ws       *websocket.Conn
	chatroom *ChatRoom
	Name     string
	in       chan *Message
}

func CreateClient(ws *websocket.Conn, chatroom *ChatRoom, name string) *Client {
	client := &Client{ws, chatroom, name, make(chan *Message)}
	go client.startAccepting()
	go client.startGenerating()
	return client
}

func (c *Client) Close() {
	c.ws.Close()
	c.chatroom.Unregister(c.Name)
}

func (c *Client) Accept(message *Message) {
	c.in <- message
}

func (c *Client) startGenerating() {
	for {
		messageType, r, _ := c.ws.NextReader()
		switch messageType {
		case websocket.CloseMessage:
			c.Close()
		case websocket.TextMessage:
			c.generate(r)
		}
	}
}

func (c *Client) generate(r io.Reader) {
	var m []byte
	r.Read(m)
	message := Message{string(m), c.Name, time.Now()}
	c.chatroom.Broadcast(&message)
}

func (c *Client) startAccepting() {
	for message := range c.in {
		m, err := json.Marshal(message)
		if err != nil {
			c.ws.WriteJSON([]byte("WHOOPS"))
		} else {
			c.ws.WriteJSON(m)
		}
	}
}

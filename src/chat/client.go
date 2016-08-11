package chat

import (
	"io"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	ws       *websocket.Conn
	chatroom *ChatRoom
	Name     string
	in       chan string
}

func CreateClient(ws *websocket.Conn, chatroom *ChatRoom, name string) *Client {
	client := &Client{ws, chatroom, name, make(chan string)}
	go client.StartMessageListener()
	go client.StartNoticationListener()
	return client
}

func (c *Client) StartMessageListener() {
	for {
		messageType, r, err := c.ws.NextReader()
		if err != nil {
			c.ws.Close()
			return
		}
		switch messageType {
		case websocket.TextMessage:
			go c.Broadcast(r)
		case websocket.BinaryMessage:
			// lol wat
			continue
		default:
			// LOL WAT
			continue
		}
	}
}

func (c *Client) StartNoticationListener() {
	for message := range c.in {
		err := c.ws.WriteJSON(message)
		if err != nil {
			// nuts
			continue
		}
	}
}

func (c *Client) Broadcast(r io.Reader) {
	var m []byte
	r.Read(m)
	message := Message{string(m), c.Name, time.Now()}
	c.chatroom.Broadcast(&message)
}

func (c *Client) Notify(obj string) {
	c.in <- obj
}

func (c *Client) Close() {
	c.ws.Close()
	c.chatroom.Unregister(c.Name)
}

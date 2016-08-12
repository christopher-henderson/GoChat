package chat

import (
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	ws           *websocket.Conn
	chatroom     *ChatRoom
	Name         string
	notifcations chan string
	incoming     chan string
}

func CreateClient(ws *websocket.Conn, chatroom *ChatRoom, name string) *Client {
	client := &Client{ws, chatroom, name, make(chan string), make(chan string)}
	go client.ListenForIncoming()
	go client.ListenForOutgoing()
	return client
}

func (c *Client) ListenForIncoming() {
	go c.Broadcast()
	for {
		messageType, r, err := c.ws.ReadMessage()
		if err != nil {
			c.ws.Close()
			return
		}
		switch messageType {
		case websocket.TextMessage:
			c.incoming <- string(r)
		case websocket.BinaryMessage:
			// lol wat
			continue
		default:
			// LOL WAT
			continue
		}
	}
}

func (c *Client) ListenForOutgoing() {
	for message := range c.notifcations {
		writer, err := c.ws.NextWriter(websocket.TextMessage)
		if err != nil {
			return
		}
		_, err = writer.Write([]byte(message))
		writer.Close()
		if err != nil {
			// nuts
			continue
		}
	}
}

func (c *Client) Broadcast() {
	for m := range c.incoming {
		message := Message{m, c.Name, time.Now()}
		c.chatroom.Broadcast(&message)
	}
}

func (c *Client) Notify(obj string) {
	c.notifcations <- obj
}

func (c *Client) Close() {
	c.ws.Close()
	c.chatroom.Unregister(c.Name)
}

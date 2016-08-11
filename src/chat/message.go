package chat

import (
	"time"
)

type Message struct {
	Content      string
	Sender       string
	TimeReceived time.Time
}

func CreateMessage(content, sender string) *Message {
	return &Message{content, sender, time.Now()}
}

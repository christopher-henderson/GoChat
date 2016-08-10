package chat

import (
	"time"
)

type Message struct {
	Content      string
	Sender       string
	TimeReceived time.Time
}

package chat

import "sync"

var MAX_SIZE int = 100

type Node struct {
	message *Message
	left    *Node
	right   *Node
}

type Log struct {
	sync.Mutex
	head *Node
	tail *Node
	size int
}

func (l *Log) Append(message *Message) {
	l.Lock()
	defer l.Unlock()
	node := Node{message, l.tail, nil}
	if l.head == nil {
		l.head = &node
	}
	if l.tail == nil {
		l.tail = &node
	} else {
		l.tail.right = &node
		l.tail = &node
	}
	l.size++
	if l.size >= MAX_SIZE {
		l.head = l.head.right
		l.head.left = nil
		l.size--
	}
}

func (l *Log) Iter() <-chan *Message {
	channel := make(chan *Message)
	go func() {
		l.Lock()
		defer l.Unlock()
		defer close(channel)
		node := l.head
		for node != nil {
			channel <- node.message
			node = node.right
		}
	}()
	return channel
}

func (l *Log) Slice() []*Message {
	var messages []*Message
	for message := range l.Iter() {
		messages = append(messages, message)
	}
	return messages
}

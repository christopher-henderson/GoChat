package chat

import "sync"

type Hub struct {
	sync.Mutex
	rooms map[string]*ChatRoom
}

func (h *Hub) Get(key string) *ChatRoom {
	h.Lock()
	defer h.Unlock()
	if _, ok := h.rooms[key]; !ok {
		h.rooms[key] = CreateChatRoom(key)
	}
	room, _ := h.rooms[key]
	return room
}

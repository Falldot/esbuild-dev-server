package ws

import "bytes"

type Hub struct {
	clients    map[*Client]bool
	Broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	reload     []byte
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		reload:     []byte("reload"),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.Broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

func (h *Hub) SendError(mes string) {
	h.Broadcast <- bytes.TrimSpace([]byte(mes))
}

func (h *Hub) SendErrorBytes(mes []byte) {
	h.Broadcast <- bytes.TrimSpace(mes)
}

func (h *Hub) SendReload() {
	h.Broadcast <- bytes.TrimSpace(h.reload)
}

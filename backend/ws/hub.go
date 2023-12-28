package ws

import (
	"encoding/json"
	"log"
)

// Hub maintains the set of active clients and broadcasts messages to the clients.
type Hub struct {
	clients    map[*Client]bool // Registered clients.
	broadcast  chan []byte      // Inbound messages from the clients.
	register   chan *Client     // Register requests from the clients.
	unregister chan *Client     // Unregister requests from clients.
}

// NewHub creates a new Hub.
func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

// Run starts the hub to process incoming register, unregister and broadcast channels.
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
		case message := <-h.broadcast:
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

func (h *Hub) broadcastScores(scores []int64) {
	scoreData, err := json.Marshal(scores)
	if err != nil {
		log.Printf("Error marshaling scores: %v", err)
		return
	}

	for client := range h.clients {
		select {
		case client.send <- scoreData:
		default:
			close(client.send)
			delete(h.clients, client)
		}
	}
}

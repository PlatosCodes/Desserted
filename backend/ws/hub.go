package ws

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/PlatosCodes/desserted/backend/util"
)

// Hub maintains the set of active clients and broadcasts messages to the clients.
type Hub struct {
	clients       map[*Client]bool  // Registered clients.
	broadcast     chan []byte       // Inbound messages from the clients.
	register      chan *Client      // Register requests from the clients.
	unregister    chan *Client      // Unregister requests from clients.
	playerClients map[int64]*Client // Maps playerGameID to Client
	messageQueue  *MessageQueue
	mutex         sync.Mutex
}

// NewHub creates a new Hub.
func NewHub() *Hub {
	mq := NewMessageQueue(100)
	return &Hub{
		broadcast:     make(chan []byte),
		register:      make(chan *Client),
		unregister:    make(chan *Client),
		clients:       make(map[*Client]bool),
		playerClients: make(map[int64]*Client),
		messageQueue:  mq,
	}
}

// Run starts the hub to process incoming register, unregister and broadcast channels.
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			h.playerClients[client.playerGameID] = client

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
					delete(h.playerClients, client.playerGameID)
					delete(h.clients, client)
				}
			}
		}
	}
}

// getClient retrieves a client by playerGameID
func (h *Hub) getClient(playerGameID int64) *Client {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	if client, ok := h.playerClients[playerGameID]; ok {
		return client
	}
	return nil
}

func (h *Hub) sendToClient(playerGameID int64, gameID int64, message interface{}) {
	client := h.getClient(playerGameID)

	if client != nil && client.gameID == gameID {
		serializedMessage, err := json.Marshal(message)
		if err != nil {
			log.Printf("Error marshaling message: %v", err)
			return
		}
		client.send <- serializedMessage
	}
}

func (h *Hub) broadcastMessage(gameID int64, message []byte) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	log.Printf("Broadcasting message to gameID %d: %s", gameID, string(message))

	for client := range h.clients {
		if client.gameID == gameID {
			select {
			case client.send <- message:
			default:
				close(client.send)
				delete(h.clients, client)
			}
		}
	}
}

func (h *Hub) broadcastExcept(exclude []int64, gameID int64, message interface{}) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	serializedMessage, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error marshaling message: %v", err)
		return
	}

	for client := range h.clients {
		if client.gameID == gameID && !util.Contains(exclude, client.playerGameID) {
			client.send <- serializedMessage
		}
	}
}

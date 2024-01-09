package ws

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	db "github.com/PlatosCodes/desserted/backend/db/sqlc"
	"github.com/gorilla/websocket"
)

const (
	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

// Define a type for message handling functions
type messageHandler func(json.RawMessage)

// Client represents a single chatting user.
type Client struct {
	ctx context.Context
	hub *Hub
	// WebSocket connection.
	conn *websocket.Conn
	// Buffered channel of outbound messages.
	send            chan []byte
	userID          int64
	store           db.Store
	messageHandlers map[string]messageHandler
	mutex           sync.Mutex
}

type WebSocketMessage struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

// NewClient initializes a new client and sets up message handlers
func NewClient(ctx context.Context, hub *Hub, conn *websocket.Conn, userID int64, store db.Store) *Client {
	client := &Client{
		ctx:             ctx,
		hub:             hub,
		conn:            conn,
		send:            make(chan []byte, 256),
		userID:          userID,
		store:           store,
		messageHandlers: make(map[string]messageHandler),
	}

	// Setting up message handlers
	client.messageHandlers["drawCard"] = client.handleDrawCard
	client.messageHandlers["playDessert"] = client.handlePlayDessert
	client.messageHandlers["endTurn"] = client.handleEndTurn
	client.messageHandlers["playSpecialCard"] = client.handlePlaySpecialCard

	return client
}

func (c *Client) handleMessage(message []byte) {
	var msg WebSocketMessage
	err := json.Unmarshal(message, &msg)
	if err != nil {
		handleWebsocketError(err, c.conn)
		return
	}

	if handler, ok := c.messageHandlers[msg.Type]; ok {
		handler(msg.Data)
	} else {
		handleWebsocketError(err, c.conn)
	}
}

// readPump pumps messages from the WebSocket connection to the Hub.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			handleWebsocketError(err, c.conn)
			break
		}
		c.handleMessage(message)
	}
}

// writePump pumps messages from the Hub to the WebSocket connection.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.mutex.Lock() // Lock for writing
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err == nil {
				w.Write(message)
				n := len(c.send)
				for i := 0; i < n; i++ {
					w.Write(<-c.send)
				}
				err = w.Close()
			}
			c.mutex.Unlock() // Unlock after writing

			if err != nil {
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			c.mutex.Lock() // Lock for writing ping
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				c.mutex.Unlock() // Unlock before returning
				return
			}
			c.mutex.Unlock() // Unlock after writing ping
		}
	}
}

// handleWebsocketError processes and logs WebSocket errors
func handleWebsocketError(err error, conn *websocket.Conn) {
	log.Printf("websocket error: %v", err)
	if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
		errMsg, _ := json.Marshal(map[string]string{"error": "WebSocket connection closed unexpectedly"})
		conn.WriteMessage(websocket.TextMessage, errMsg)
	}
}

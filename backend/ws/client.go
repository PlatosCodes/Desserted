package ws

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	db "github.com/PlatosCodes/desserted/backend/db/sqlc"
	gameservice "github.com/PlatosCodes/desserted/backend/game_service"
	"github.com/gorilla/websocket"
)

const (
	// Time allowed to read the next pong message from the peer.
	pongWait = 5 * time.Minute

	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 1024
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
	gameService     *gameservice.GameService
	messageHandlers map[string]messageHandler
	messageQueue    *MessageQueue
	broadcastChan   chan<- []byte
	mutex           sync.Mutex
	gameID          int64
	playerGameID    int64
}

type WebSocketMessage struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

// NewClient initializes a new client and sets up message handlers
func NewClient(ctx context.Context, hub *Hub, conn *websocket.Conn, userID int64, store db.Store, gameService *gameservice.GameService, mq *MessageQueue, broadcastChan chan<- []byte, gameID int64, playerGameID int64) *Client {
	client := &Client{
		ctx:             ctx,
		hub:             hub,
		conn:            conn,
		send:            make(chan []byte, 256),
		userID:          userID,
		store:           store,
		gameService:     gameService,
		messageHandlers: make(map[string]messageHandler),
		messageQueue:    mq,
		broadcastChan:   broadcastChan,
		gameID:          gameID,
		playerGameID:    playerGameID,
	}

	// Setting up message handlers
	client.messageHandlers["drawCard"] = client.handleDrawCard
	client.messageHandlers["playDessert"] = client.handlePlayDessert
	client.messageHandlers["endTurn"] = client.handleEndTurn
	client.messageHandlers["playSpecialCard"] = client.handlePlaySpecialCard

	return client
}

// func (c *Client) sendBroadcastMessage(msg []byte) {
// 	c.broadcastChan <- msg
// }

func (c *Client) handleMessage(message []byte) {
	var msg WebSocketMessage

	log.Printf("Received message: %s", string(message))

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

	if err := c.conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Printf("SetReadDeadline failed: %v", err)
		return
	}

	c.conn.SetPongHandler(func(string) error {
		return c.conn.SetReadDeadline(time.Now().Add(pongWait))
	})

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
			if err := c.conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				log.Printf("SetWriteDeadline failed: %v", err)
				return
			}

			if !ok {
				if err := c.conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					log.Printf("WriteMessage failed: %v", err)
				}
				return
			}

			c.mutex.Lock()
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				c.mutex.Unlock()
				log.Printf("NextWriter failed: %v", err)
				return
			}

			if _, err := w.Write(message); err != nil {
				log.Printf("Write failed: %v", err)
			}

			if err := w.Close(); err != nil {
				log.Printf("Close failed: %v", err)
			}
			c.mutex.Unlock()

		case <-ticker.C:
			if err := c.conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				log.Printf("SetWriteDeadline failed: %v", err)
				return
			}
			c.mutex.Lock()
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Printf("WriteMessage failed: %v", err)
				c.mutex.Unlock()
				return
			}
			c.mutex.Unlock()
		}
	}
}

// handleWebsocketError processes and logs WebSocket errors
func handleWebsocketError(err error, conn *websocket.Conn) {
	log.Printf("websocket error: %v", err)
	if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
		errMsg, _ := json.Marshal(map[string]string{"error": "WebSocket connection closed unexpectedly"})
		if err := conn.WriteMessage(websocket.TextMessage, nil); err != nil {
			log.Printf("WriteMessage failed: %v", errMsg)
			return
		}
	}
}

func (c *Client) sendErrorMessage(errorMessage string) {
	errorResponse := struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	}{
		Type:    "error",
		Message: errorMessage,
	}

	msg, err := json.Marshal(errorResponse)
	if err != nil {
		log.Printf("Error marshaling error message: %v", err)
		return
	}

	c.send <- msg
}

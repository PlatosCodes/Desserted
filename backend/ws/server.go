package ws

import (
	"context"
	"fmt"
	"net/http"

	db "github.com/PlatosCodes/desserted/backend/db/sqlc"
	"github.com/PlatosCodes/desserted/backend/token"
	"github.com/PlatosCodes/desserted/backend/util"

	"github.com/gorilla/websocket"
)

func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request, config util.Config, store db.Store, tokenMaker token.Maker) {
	// Define upgrader here with dynamic CheckOrigin
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return r.Header.Get("Origin") == config.FrontendAddress
		},
	}

	// Extract token from query parameters
	accessToken := r.URL.Query().Get("access_token")
	if accessToken == "" {
		http.Error(w, "Authorization token is required", http.StatusUnauthorized)
		return
	}

	// Verify the token
	payload, err := validateToken(r, tokenMaker)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		handleWebsocketError(err, conn)
		return
	}

	// Create a context for the client
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client := NewClient(ctx, hub, conn, payload.UserID, store)

	client.hub.register <- client

	go client.writePump()
	go client.readPump()
}

// validateToken extracts and validates the token from the request.
func validateToken(r *http.Request, tokenMaker token.Maker) (*token.Payload, error) {
	accessToken := r.URL.Query().Get("access_token")
	if accessToken == "" {
		return nil, fmt.Errorf("authorization token is required")
	}
	return tokenMaker.VerifyToken(accessToken)
}

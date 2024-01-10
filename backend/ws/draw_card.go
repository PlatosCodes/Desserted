package ws

import (
	"context"
	"encoding/json"
	"log"

	db "github.com/PlatosCodes/desserted/backend/db/sqlc"
	pb "github.com/PlatosCodes/desserted/backend/pb"

	"github.com/gorilla/websocket"
)

// DrawCardPayload matches the structure expected in the DrawCardRequest message
type DrawCardPayload struct {
	GameID       int64 `json:"game_id"`
	PlayerGameID int64 `json:"player_game_id"`
	PlayerNumber int32 `json:"player_number"`
}

func (c *Client) handleDrawCard(payload json.RawMessage) {
	// Unmarshal the payload into the expected structure
	var drawCardPayload DrawCardPayload
	if err := json.Unmarshal(payload, &drawCardPayload); err != nil {
		log.Printf("Error unmarshaling draw card payload: %v", err)
		sendErrorMessage(c.conn, "Invalid payload for drawing card")
		return
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Call the DrawCard function from your gRPC service
	card, err := c.store.DrawCard(context.Background(), db.DrawCardTxParams{
		GameID:       drawCardPayload.GameID,
		PlayerID:     drawCardPayload.PlayerGameID,
		PlayerNumber: drawCardPayload.PlayerNumber,
	})
	if err != nil {
		log.Printf("Error drawing card: %v", err)
		sendErrorMessage(c.conn, "Error drawing card from the deck")
		return
	}

	// // Create a response message
	// response := pb.DrawCardResponse{
	// 	CardId: cardID,
	// }

	// Send the response back to the client
	// sendDrawCardResponse(c.conn, &response)

	drawCardUpdateMessage := prepareDrawCardUpdateMessage(card)

	// msg, err := json.Marshal(drawCardUpdateMessage)
	// if err != nil {
	// 	log.Printf("Error marshaling draw card response: %v", err)
	// 	return
	// }

	// Send the message through the WebSocket connection
	if err := c.conn.WriteMessage(websocket.TextMessage, drawCardUpdateMessage); err != nil {
		log.Printf("Error sending draw card response: %v", err)
	}
	// c.hub.broadcast <- drawCardUpdateMessage

}

// sendErrorMessage sends an error message to the client
func sendErrorMessage(conn *websocket.Conn, errorMessage string) {
	// Implement a function to send an error message to the client
	// This can be a simple JSON message with an 'error' field
}

// sendDrawCardResponse sends a draw card response to the client
func sendDrawCardResponse(conn *websocket.Conn, response *pb.DrawCardResponse) {
	// Marshal the response into JSON
	msg, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshaling draw card response: %v", err)
		return
	}

	// Send the message through the WebSocket connection
	if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		log.Printf("Error sending draw card response: %v", err)
	}
}

func prepareDrawCardUpdateMessage(card db.DrawCardTxResult) []byte {
	// Define a struct for the message
	type DrawCardUpdate struct {
		Type string              `json:"type"`
		Card db.DrawCardTxResult `json:"card"`
	}

	updateMsg := DrawCardUpdate{
		Type: "drawCardResponse",
		Card: card,
	}

	msg, err := json.Marshal(updateMsg)
	if err != nil {
		log.Printf("Error marshaling score update message: %v", err)
		return nil
	}

	return msg
}

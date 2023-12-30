package ws

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"

	db "github.com/PlatosCodes/desserted/backend/db/sqlc"
	pb "github.com/PlatosCodes/desserted/backend/pb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/gorilla/websocket"
)

type EndTurnPayload struct {
	GameID       int64 `json:"game_id"`
	PlayerGameID int64 `json:"player_game_id"`
}

func (c *Client) handleEndTurn(payload json.RawMessage) {
	// Unmarshal the payload into the expected structure
	var endTurnPayload EndTurnPayload
	if err := json.Unmarshal(payload, &endTurnPayload); err != nil {
		log.Printf("Error unmarshaling draw card payload: %v", err)
		sendErrorMessage(c.conn, "Invalid payload for drawing card")
		return
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	ctx := context.Background()

	game, err := c.store.GetGameByID(ctx, endTurnPayload.GameID)
	if err != nil {
		log.Printf("Error getting game info: %v", err)
		sendErrorMessage(c.conn, "Error getting game info")
		return
	}

	// Call the UpdatePlayerNumber function from gRPC service
	err = c.store.UpdateGameState(ctx, db.UpdateGameStateParams{
		GameID:      endTurnPayload.GameID,
		CurrentTurn: game.CurrentTurn + 1,
		CurrentPlayerNumber: sql.NullInt32{
			Int32: ((game.CurrentPlayerNumber.Int32 + int32(1)) % game.NumberOfPlayers) + 1,
			Valid: true,
		},
	})
	if err != nil {
		log.Printf("Error updating game state : %v", err)
		sendErrorMessage(c.conn, "Error updating game state")
		return
	}

	updatedGame, err := c.store.GetGameByID(ctx, endTurnPayload.GameID)
	if err != nil {
		log.Printf("Error getting game info: %v", err)
		sendErrorMessage(c.conn, "Error getting game info")
		return
	}

	// Create a response message
	response := pb.GetGameByIDResponse{
		Game: &pb.Game{
			GameId:              updatedGame.GameID,
			Status:              updatedGame.Status,
			CreatedBy:           updatedGame.CreatedBy,
			NumberOfPlayers:     updatedGame.NumberOfPlayers,
			CurrentTurn:         updatedGame.CurrentTurn,
			CurrentPlayerNumber: updatedGame.CurrentPlayerNumber.Int32,
			StartTime:           timestamppb.New(updatedGame.StartTime),
			EndTime:             nil,
		},
	}

	// Send the response back to the client
	sendEndTurnResponse(c.conn, &response)
}

// // sendErrorMessage sends an error message to the client
// func sendErrorMessage(conn *websocket.Conn, errorMessage string) {
// 	// Implement a function to send an error message to the client
// 	// This can be a simple JSON message with an 'error' field
// }

// sendEndTurnResponse sends a game response to the client
func sendEndTurnResponse(conn *websocket.Conn, response *pb.GetGameByIDResponse) {
	// Marshal the response into JSON
	msg, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshaling draw card response: %v", err)
		return
	}

	// Send the message through the WebSocket connection
	if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		log.Printf("Error sending end turn response: %v", err)
	}
}

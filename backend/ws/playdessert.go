package ws

import (
	"context"
	"encoding/json"
	"log"

	db "github.com/PlatosCodes/desserted/backend/db/sqlc"
	"github.com/gorilla/websocket"
)

// PlayDessertPayload mirrors the PlayDessertRequest from gRPC
type PlayDessertPayload struct {
	PlayerGameID int64   `json:"player_game_id"`
	DessertName  string  `json:"dessert_name"`
	CardIDs      []int64 `json:"card_ids"`
}

// Custom struct that mirrors pb.PlayDessertResponse but is safe for marshaling
type PlayDessertResponseJSON struct {
	DessertPlayedId int64           `json:"dessert_played_id"`
	PlayerGame      *PlayerGameJSON `json:"player_game"`
	GameOver        bool            `json:"game_over"`
}

type PlayerGameJSON struct {
	PlayerGame   int64  `json:"player_game_id"`
	PlayerId     int64  `json:"player_id"`
	GameId       int64  `json:"game_id"`
	PlayerScore  int32  `json:"player_score"`
	PlayerStatus string `json:"player_status"`
}

func (c *Client) handlePlayDessert(payload json.RawMessage) {
	var playDessertPayload PlayDessertPayload
	if err := json.Unmarshal(payload, &playDessertPayload); err != nil {
		log.Printf("Error unmarshaling play dessert payload: %v", err)
		// Ideally, send an error message back to the client
		return
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	ctx := context.Background()

	result, err := c.store.PlayDessertTx(ctx, db.PlayDessertTxParams{
		PlayerGameID: playDessertPayload.PlayerGameID,
		DessertName:  playDessertPayload.DessertName,
		CardIDs:      playDessertPayload.CardIDs,
	})

	if err != nil {
		log.Println("and the error is:", err)
		// Handle error, send response back to client
		return
	}

	// Convert protobuf response to custom JSON struct
	responseJSON := PlayDessertResponseJSON{
		DessertPlayedId: result.DessertPlayedID,
		PlayerGame: &PlayerGameJSON{
			PlayerGame:   result.PlayerGame.PlayerGameID,
			PlayerId:     result.PlayerGame.PlayerID,
			GameId:       result.PlayerGame.GameID,
			PlayerScore:  result.PlayerGame.PlayerScore.Int32,
			PlayerStatus: result.PlayerGame.PlayerStatus.String,
		},
		GameOver: result.GameOver,
	}

	// Now marshal and send this custom struct
	data, err := json.Marshal(responseJSON)
	if err != nil {
		log.Printf("Error marshaling response: %v", err)
		// Send error back to client
		return
	}
	c.conn.WriteMessage(websocket.TextMessage, data)
}

package ws

import (
	"context"
	"encoding/json"
	"log"

	gameservice "github.com/PlatosCodes/desserted/backend/game_service"
)

// PlayDessertPayload mirrors the PlayDessertRequest from gRPC
type PlayDessertPayload struct {
	GameID       int64   `json:"game_id"`
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
	PlayerNumber int32  `json:"player_number"`
	PlayerScore  int32  `json:"player_score"`
	PlayerStatus string `json:"player_status"`
}

func (c *Client) handlePlayDessert(payload json.RawMessage) {
	var playDessertPayload PlayDessertPayload
	if err := json.Unmarshal(payload, &playDessertPayload); err != nil {
		log.Printf("Error unmarshaling play dessert payload: %v", err)
		c.sendErrorMessage(err.Error())
		return
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	ctx := context.Background()

	// Forward playDessertPayload to the GameService's PlayDessertHandler
	err := c.gameService.PlayDessertHandler(ctx, gameservice.PlayDessertHandlerParams{
		PlayerGameID: playDessertPayload.PlayerGameID,
		GameID:       playDessertPayload.GameID,
		DessertName:  playDessertPayload.DessertName,
		CardIDs:      playDessertPayload.CardIDs,
	})
	if err != nil {
		log.Println("Error processing play dessert event:", err)
		c.sendErrorMessage(err.Error())
		return
	}

}

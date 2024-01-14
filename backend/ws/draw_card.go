package ws

import (
	"context"
	"encoding/json"
	"log"

	gameservice "github.com/PlatosCodes/desserted/backend/game_service"
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
		c.sendErrorMessage(err.Error())
		return
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	ctx := context.Background()
	// Call the DrawCardHandler function from game service
	err := c.gameService.DrawCardHandler(ctx, gameservice.DrawCardHandlerParams{
		GameID:       drawCardPayload.GameID,
		PlayerGameID: drawCardPayload.PlayerGameID,
		PlayerNumber: drawCardPayload.PlayerNumber,
	})
	if err != nil {
		log.Printf("Error drawing card: %v", err)
		c.sendErrorMessage(err.Error())
		return
	}

}

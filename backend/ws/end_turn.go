package ws

import (
	"context"
	"encoding/json"
	"log"
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

	err := c.gameService.EndTurnHandler(ctx, endTurnPayload.GameID, endTurnPayload.PlayerGameID)
	if err != nil {
		log.Println("Error processing end turn event:", err)
		c.sendErrorMessage(err.Error())
		return
	}

}

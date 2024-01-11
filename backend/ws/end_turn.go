package ws

import (
	"context"
	"encoding/json"
	"log"

	pb "github.com/PlatosCodes/desserted/backend/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type EndTurnPayload struct {
	GameID       int64 `json:"game_id"`
	PlayerGameID int64 `json:"player_game_id"`
}

type GameJSON struct {
	GameID              int64  `json:"game_id"`
	Status              string `json:"status"`
	CreatedBy           int64  `json:"created_by"`
	NumberOfPlayers     int32  `json:"number_of_players"`
	CurrentTurn         int32  `json:"current_turn"`
	CurrentPlayerNumber int32  `json:"current_player_number"`
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

	updatedGame, err := c.store.EndTurnTx(ctx, endTurnPayload.GameID, endTurnPayload.PlayerGameID)
	if err != nil {
		log.Printf("Error ending turn: %v", err)
		sendErrorMessage(c.conn, "Error ending turn")
		return
	}

	updatedPbGame := &pb.Game{
		GameId:              updatedGame.Game.GameID,
		Status:              updatedGame.Game.Status,
		CreatedBy:           updatedGame.Game.CreatedBy,
		NumberOfPlayers:     updatedGame.Game.NumberOfPlayers,
		CurrentTurn:         updatedGame.Game.CurrentTurn,
		CurrentPlayerNumber: updatedGame.Game.CurrentPlayerNumber.Int32,
		StartTime:           timestamppb.New(updatedGame.Game.StartTime),
		EndTime:             nil,
	}

	endTurnUpdateMsg := prepareEndTurnUpdateMessage(updatedPbGame)
	c.messageQueue.Enqueue(endTurnUpdateMsg)

}

func prepareEndTurnUpdateMessage(game *pb.Game) []byte {
	// Define a struct for the message
	type EndTurnUpdate struct {
		Type string   `json:"type"`
		Game GameJSON `json:"game"`
	}

	gameJSON := GameJSON{
		GameID:              game.GameId,
		Status:              game.Status,
		CreatedBy:           game.CreatedBy,
		NumberOfPlayers:     game.NumberOfPlayers,
		CurrentTurn:         game.CurrentTurn,
		CurrentPlayerNumber: game.CurrentPlayerNumber,
	}

	updateMsg := EndTurnUpdate{
		Type: "endTurn",
		Game: gameJSON,
	}

	msg, err := json.Marshal(updateMsg)
	if err != nil {
		log.Printf("Error marshaling score update message: %v", err)
		return nil
	}

	return msg
}

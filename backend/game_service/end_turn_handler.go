package gameservice

import (
	"context"
	"fmt"
	"log"
)

// Handles the coordination of services when a player's turn is over
func (s *GameService) EndTurnHandler(ctx context.Context, gameID int64, playerGameID int64) error {

	updatedGame, err := s.store.EndTurnTx(ctx, gameID, playerGameID)
	if err != nil {
		log.Printf("Error ending turn: %v", err)
		return fmt.Errorf("error ending turn: %v", err)
	}

	endTurnData := EndTurnData{
		GameID:              updatedGame.Game.GameID,
		Status:              updatedGame.Game.Status,
		CreatedBy:           updatedGame.Game.CreatedBy,
		NumberOfPlayers:     updatedGame.Game.NumberOfPlayers,
		CurrentTurn:         updatedGame.Game.CurrentTurn,
		CurrentPlayerNumber: updatedGame.Game.CurrentPlayerNumber.Int32,
	}

	updateMsg := Event{
		Type: EventTypeEndTurn,
		Data: endTurnData,
	}

	EmitEvent(updateMsg)

	// Reset the player's actions taken for next turn
	err = s.store.ResetTurnActions(ctx, playerGameID)
	if err != nil {
		log.Printf("Error resetting playr's actions for next turn turn: %v", err)
		return fmt.Errorf("error resetting playr's actions for next turn turn: %v", err)
	}

	return nil

}

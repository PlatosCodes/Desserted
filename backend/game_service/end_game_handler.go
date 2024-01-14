package gameservice

import (
	"context"
	"fmt"
	"log"
)

// Handles the coordination of services the game is over
func (s *GameService) EndGameHandler(ctx context.Context, gameID int64, playerGameID int64) error {

	// db end game, declare winner functions

	err := s.store.EndGame(ctx, gameID)
	if err != nil {
		log.Printf("Error ending game: %v", err)
		return fmt.Errorf("error ending game: %v", err)
	}

	game, err := s.store.GetGameByID(ctx, gameID)
	if err != nil {
		log.Printf("Error getting game info to end game: %v", err)
		return fmt.Errorf("error getting game info to end game: %v", err)
	}

	winner, err := s.store.DeclareWinner(ctx, gameID)
	if err != nil {
		log.Printf("Error getting game winner info: %v", err)
		return fmt.Errorf("error getting game winner info: %v", err)
	}

	endGameData := EndGameData{
		GameID:              gameID,
		Status:              game.Status,
		WinningPlayerGameID: winner.PlayerGameID,
		WinningPlayerNumber: winner.PlayerNumber.Int32,
		WinningScore:        winner.PlayerScore,
	}

	updateMsg := Event{
		Type: EventTypeEndGame,
		Data: endGameData,
	}

	EmitEvent(updateMsg)

	return nil

}

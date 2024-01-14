package gameservice

import (
	"context"
	"errors"
	"fmt"
	"log"

	db "github.com/PlatosCodes/desserted/backend/db/sqlc"
)

type DrawCardHandlerParams struct {
	GameID       int64 `json:"game_id"`
	PlayerGameID int64 `json:"player_game_id"`
	PlayerNumber int32 `json:"player_number"`
}

// Handles the coordination of services when a player draws a card
func (s *GameService) DrawCardHandler(ctx context.Context, arg DrawCardHandlerParams) error {

	// Check if it's the player's turn
	game, err := s.store.GetGameByID(ctx, arg.GameID)
	if err != nil {
		return err
	}
	if game.CurrentPlayerNumber.Int32 != arg.PlayerNumber {
		return errors.New("not the player's turn")
	}

	// Check if player has already drawn a card this turn
	hasAlreadyDrawnCard, err := s.store.CheckCardDrawn(ctx, arg.PlayerGameID)
	if err != nil {
		return fmt.Errorf("failed to get player's drawn card status for this turn: %w", err)
	}

	if hasAlreadyDrawnCard {
		return fmt.Errorf("player has already drawn card this turn: %w", err)
	}

	card, err := s.store.DrawCard(ctx, db.DrawCardTxParams{
		GameID:       arg.GameID,
		PlayerID:     arg.PlayerGameID,
		PlayerNumber: arg.PlayerNumber,
	})
	if err != nil {
		log.Printf("Error drawing card game: %v", err)
		return fmt.Errorf("error drawing card: %v", err)
	}

	//Check if all actions are completed
	completed, err := s.store.CheckAllActionsCompleted(ctx, arg.PlayerGameID)
	if err != nil {
		log.Printf("Error checking actions completed: %v", err)
		return fmt.Errorf("error checking actions completed: %v", err)

	}

	if completed.Bool {
		log.Println("player's turn is completed -- they have drawn card and played both a dessert and special card")
		s.EndTurnHandler(ctx, arg.GameID, card.PlayerGameID)
	}

	drawCardEvent := Event{
		Type: EventTypeCardDrawn,
		Data: CardDrawnData{
			CardID:       card.CardID,
			CardName:     card.CardName,
			PlayerGameID: card.PlayerGameID,
			PlayerHandID: card.PlayerHandID,
			GameID:       arg.GameID,
		},
	}

	EmitEvent(drawCardEvent)

	return nil
}

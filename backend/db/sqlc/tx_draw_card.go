package db

import (
	"context"
	"database/sql"
	"fmt"
)

// DrawTxParams holds parameters for the StartGameTx function
type DrawCardTxParams struct {
	GameID       int64 `json:"game_id"`
	PlayerID     int64 `json:"player_id"`
	PlayerNumber int32 `json:"player_number"`
}

type DrawCardTxResult struct {
	CardID       int64  `json:"card_id"`
	CardName     string `json:"name"`
	PlayerGameID int64  `json:"player_game_id"`
	PlayerHandID int64  `json:"player_hand_id"`
}

// DrawCard draws the top card from the deck for a given game.
func (store *SQLStore) DrawCard(ctx context.Context, arg DrawCardTxParams) (DrawCardTxResult, error) {
	var cardID int64

	var rsp DrawCardTxResult

	// Begin transaction
	tx, err := store.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return rsp, fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Draw the top card from the game deck
	cardID, err = store.DrawTopCard(ctx, arg.GameID)
	if err != nil {
		return rsp, fmt.Errorf("failed to draw top card: %w", err)
	}

	player_hand_id, err := store.AddCardToPlayerHand(ctx, AddCardToPlayerHandParams{
		PlayerGameID: arg.PlayerID,
		CardID:       cardID,
	})
	if err != nil {
		return rsp, fmt.Errorf("failed to add card to player's hand: %w", err)
	}

	// Remove the drawn card from the game deck
	err = store.RemoveCardFromDeck(ctx, RemoveCardFromDeckParams{
		GameID: arg.GameID,
		CardID: cardID,
	})
	if err != nil {
		return rsp, fmt.Errorf("failed to remove card from game deck: %w", err)
	}

	card, err := store.GetCardByID(ctx, cardID)
	if err != nil {
		return rsp, fmt.Errorf("failed to get card info from database: %w", err)
	}

	rsp = DrawCardTxResult{
		CardID:       cardID,
		CardName:     card.Name,
		PlayerGameID: arg.PlayerID,
		PlayerHandID: player_hand_id,
	}

	err = store.UpdateCardDrawnStatus(ctx, arg.PlayerID)
	if err != nil {
		return rsp, fmt.Errorf("failed to update player's drawn card status for this turn: %w", err)
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return rsp, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return rsp, nil
}

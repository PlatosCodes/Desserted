package db

import (
	"context"
	"fmt"
	"log"
)

type PlayDessertTxParams struct {
	PlayerGameID    int64
	DessertName     string
	Cards           []Card
	SpecialCardUsed bool
	Score           int32
}

// Handles the database transactions related to playing a dessert
func (store *SQLStore) PlayDessertTx(ctx context.Context, arg PlayDessertTxParams) (PlayerGame, error) {
	var updatedPlayerGame PlayerGame

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Record each card played and remove from hand
		for _, card := range arg.Cards {
			err = q.RecordPlayedCard(ctx, RecordPlayedCardParams{
				PlayerGameID: arg.PlayerGameID,
				CardID:       card.CardID,
			})
			if err != nil {
				return fmt.Errorf("error recording card played: %w", err)
			}

			err = q.RemoveCardFromPlayerHand(ctx, RemoveCardFromPlayerHandParams{
				PlayerGameID: arg.PlayerGameID,
				CardID:       card.CardID,
			})
			if err != nil {
				return fmt.Errorf("error removing card from hand: %w", err)
			}
		}

		// Get the DessertID
		dessert, err := q.GetDessertByName(ctx, arg.DessertName)
		if err != nil {
			return fmt.Errorf("error fetching dessert: %w", err)
		}

		// Record the Dessert being played
		err = q.RecordDessertPlayed(ctx, RecordDessertPlayedParams{
			PlayerGameID: arg.PlayerGameID,
			DessertID:    dessert.DessertID,
		})
		if err != nil {
			return fmt.Errorf("error recording dessert played: %w", err)
		}

		// Update dessert played status
		err = q.UpdateDessertPlayedStatus(ctx, arg.PlayerGameID)
		if err != nil {
			log.Printf("Error updating dessert played status: %v", err)
		}

		// Get Player's current game info to access their current score
		currPlayer, err := q.GetPlayerGame(ctx, arg.PlayerGameID)
		if err != nil {
			return fmt.Errorf("error getting player's previous score: %w", err)
		}

		//Update player's score
		updatedPlayerGame, err = q.UpdatePlayerScore(ctx, UpdatePlayerScoreParams{
			PlayerGameID: arg.PlayerGameID,
			PlayerScore:  currPlayer.PlayerScore + arg.Score,
		})
		if err != nil {
			return fmt.Errorf("error updating player's score: %w", err)
		}

		// Set Dessert Played for this turn to true
		err = q.UpdateDessertPlayedStatus(ctx, arg.PlayerGameID)
		if err != nil {
			return fmt.Errorf("error updating player's dessert played status: %w", err)
		}

		return nil
	})

	return updatedPlayerGame, err
}

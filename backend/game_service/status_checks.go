package gameservice

import (
	"context"
	"errors"
	"fmt"

	db "github.com/PlatosCodes/desserted/backend/db/sqlc"
	"github.com/PlatosCodes/desserted/backend/game"
)

// Check if it's the player's turn
func (s *GameService) checkPlayerTurn(ctx context.Context, gameID int64, playerGameID int64) error {
	game, err := s.store.GetGameByID(ctx, gameID)
	if err != nil {
		return fmt.Errorf("error getting game id: %w", err)
	}

	playerGame, err := s.store.GetPlayerGame(ctx, playerGameID)
	if err != nil {
		return fmt.Errorf("error getting player game: %w", err)
	}

	if game.CurrentPlayerNumber.Int32 != playerGame.PlayerNumber.Int32 {
		return errors.New("not the player's turn")
	}

	return nil
}

func (s *GameService) checkAlreadyPlayedDessert(ctx context.Context, playerGameID int64) error {
	// Query the database to check if the player has already played a dessert
	alreadyPlayed, err := s.store.CheckDessertPlayed(ctx, playerGameID)
	if err != nil {
		return fmt.Errorf("error checking if player has played a dessert this turn: %w", err)
	}

	if alreadyPlayed {
		// Return an error or some form of indication that the player has already played a dessert
		return fmt.Errorf("player has already played a dessert this turn: %w", err)

	}

	return nil
}

// func CheckGameWinningCondition(ctx context.Context, playerGameID int64, q *db.Queries) (bool, error) {
// 	// Logic to check if the game is won
// 	// ...

// 	return isGameWon, nil
// }

// ValidateCardsInHand ensures that each card is valid, in the player's hand, and includes no more than one special card.
func validateCardsInHand(ctx context.Context, store db.Store, playerGameID int64, cardIDs []int64) ([]db.Card, *db.Card, error) {
	var ingredientsList []db.Card
	var specialCard *db.Card

	for _, cardID := range cardIDs {
		card, err := store.GetCardByID(ctx, cardID)
		if err != nil {
			return nil, nil, fmt.Errorf("error fetching card: %w", err)
		}

		inHand, err := store.IsCardInPlayerHand(ctx, db.IsCardInPlayerHandParams{
			PlayerGameID: playerGameID,
			CardID:       cardID,
		})
		if err != nil || !inHand {
			return nil, nil, fmt.Errorf("card not in player's hand: %w", err)
		}

		if card.Type == game.Special {

			if specialCard != nil {
				return nil, nil, fmt.Errorf("hand cannot contain more than one special card: %w", err)
			}
			specialCard = &db.Card{
				CardID: card.CardID,
				Type:   card.Type,
				Name:   card.Name,
			}
			continue
		}

		ingredientsList = append(ingredientsList, card)
	}

	return ingredientsList, specialCard, nil
}

// // Validate and process each card
// for _, cardID := range arg.CardIDs {
// 	// Fetch each card and validate
// 	_, err := q.GetCardByID(ctx, cardID)
// 	if err != nil {
// 		return fmt.Errorf("error fetching card: %w", err)
// 	}

// 	// Check if card is in player's hand
// 	inHand, err := q.IsCardInPlayerHand(ctx, IsCardInPlayerHandParams{
// 		PlayerGameID: arg.PlayerGameID,
// 		CardID:       cardID,
// 	})
// 	if err != nil || !inHand {
// 		return fmt.Errorf("card not in player's hand: %w", err)
// 	}

// 	// Record card played and remove from player's hand
// 	err = q.RecordPlayedCard(ctx, RecordPlayedCardParams{
// 		PlayerGameID: arg.PlayerGameID,
// 		CardID:       cardID,
// 	})
// 	if err != nil {
// 		return fmt.Errorf("error recording card played: %w", err)
// 	}

// 	err = q.RemoveCardFromPlayerHand(ctx, RemoveCardFromPlayerHandParams{
// 		PlayerGameID: arg.PlayerGameID,
// 		CardID:       cardID,
// 	})
// 	if err != nil {
// 		return fmt.Errorf("error removing card from hand: %w", err)
// 	}

// 	if card.Type == util.Special {

// 		// Check if player has already played a special card this turn
// 		hasAlreadyPlayedSpecialCard, err := store.CheckSpecialCardPlayed(ctx, arg.PlayerGameID)
// 		if err != nil {
// 			return fmt.Errorf("failed to get player's drawn card status for this turn: %w", err)
// 		}

// 		if hasAlreadyPlayedSpecialCard {
// 			return fmt.Errorf("player has already played a special card this turn: %w", err)
// 		}

// 		switch card.Name {
// 		case "Wildcard Ingredient":
// 			wildcardUsed = true
// 		case "Double Points":
// 			doublePointsMultiplier = 2
// 		case "Glass of Milk":
// 			extraPoints += 3
// 		case "Mystery Ingredient":
// 			extraPoints += util.RandomPoints()
// 		}
// 		// Update special card played status
// 		err = q.UpdateSpecialCardPlayedStatus(ctx, arg.PlayerGameID)
// 		if err != nil {
// 			log.Printf("Error updating special card played status: %v", err)
// 		}
// 		continue // Skip adding special cards to ingredientsList
// 	}

// 	ingredientsList = append(ingredientsList, card.Name)
// }

// if wildcardUsed {
// 	wildcardUsed = false
// }

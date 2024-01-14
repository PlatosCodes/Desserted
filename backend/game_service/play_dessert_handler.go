package gameservice

import (
	"context"
	"errors"
	"fmt"
	"log"

	db "github.com/PlatosCodes/desserted/backend/db/sqlc"
	"github.com/PlatosCodes/desserted/backend/game"
)

type PlayDessertHandlerParams struct {
	GameID       int64   `json:"game_id"`
	PlayerGameID int64   `json:"player_game_id"`
	DessertName  string  `json:"dessert_name"`
	CardIDs      []int64 `json:"card_ids"`
}

// Handles the coordination of services when a player play's a dessert
func (s *GameService) PlayDessertHandler(ctx context.Context, arg PlayDessertHandlerParams) error {

	// Perform checks and validations
	// Check if it is the player's turn
	if err := s.checkPlayerTurn(ctx, arg.GameID, arg.PlayerGameID); err != nil {
		return err
	}

	// Check if player has already played a dessert this turn
	if err := s.checkAlreadyPlayedDessert(ctx, arg.PlayerGameID); err != nil {
		return err
	}

	// Check if player is playing a valid dessert
	if !game.IsSupportedDessertType(arg.DessertName) {
		return errors.New("error checking actions completed: Invalid dessert")
	}

	// Validate cards in player's hand
	ingredientsList, specialCard, err := validateCardsInHand(ctx, s.store, arg.PlayerGameID, arg.CardIDs)
	if err != nil {
		return err
	}

	// Get the required ingredients for the dessert from the game package
	requiredIngredients, err := game.GetRequiredIngredientsForDessert(arg.DessertName)
	if err != nil {
		return err
	}

	// Check if wildcard ingredient is being played and, if so, add the missing ingredient to the ingredient list
	if specialCard != nil && specialCard.Name == "Wildcard Ingredient" {
		ingredientsList, err = s.appendMissingIngredientWithWildcard(ctx, ingredientsList, requiredIngredients)
		if err != nil {
			return err
		}
	}

	// convert ingredient list to ingredient names for dessert validation by game package
	var ingredientNames []string
	for _, ingredient := range ingredientsList {
		ingredientNames = append(ingredientNames, ingredient.Name)
	}

	// Validate if the player's ingredient list matches the required ingredients for the dessert played
	if err = game.ValidateDessert(arg.DessertName, ingredientNames); err != nil {
		return err
	}

	// Calculate the score for the played dessert and special card bonuses
	var specialCardName string
	if specialCard != nil {
		specialCardName = specialCard.Name
	}
	score, err := game.CalculateDessertScore(arg.DessertName, ingredientNames, specialCardName)
	if err != nil {
		return err
	}

	// All validations passed, proceed to PlayDessertTx
	updatedPlayerGame, err := s.store.PlayDessertTx(ctx, db.PlayDessertTxParams{
		PlayerGameID:    arg.PlayerGameID,
		DessertName:     arg.DessertName,
		Cards:           ingredientsList,
		SpecialCardUsed: specialCard != nil,
		Score:           score,
	})
	if err != nil {
		return err
	}

	// If specialCard != nil, update Special Card played in the database
	if specialCard != nil {
		s.store.UpdateSpecialCardPlayedStatus(ctx, arg.PlayerGameID)
	}

	// Check if the game is won
	winningCondition := game.IsGameWon(updatedPlayerGame.PlayerScore)

	dessertPlayedEvent := Event{
		Type: EventTypeDessertPlayed,
		Data: DessertPlayedData{
			GameID:       arg.GameID,
			PlayerGameID: arg.PlayerGameID,
			PlayerNumber: updatedPlayerGame.PlayerNumber.Int32,
			DessertName:  arg.DessertName,
			Score:        updatedPlayerGame.PlayerScore,
			DessertScore: score,
			GameOver:     winningCondition,
			Success:      true,
		},
	}

	EmitEvent(dessertPlayedEvent)

	// Check if all player actions for this turn are completed
	completed, err := s.store.CheckAllActionsCompleted(ctx, arg.PlayerGameID)
	if err != nil {
		log.Printf("Error checking actions completed: %v", err)
		return fmt.Errorf("error checking actions completed: %v", err)
	}

	if completed.Bool {
		// notify EndTurnHandler to send event message to end turn, and reset the player's turn actions for next turn
		s.EndTurnHandler(ctx, arg.GameID, arg.PlayerGameID)
		log.Printf("Player %v has completed all actions for this turn", updatedPlayerGame.PlayerNumber)
	}

	err = s.UpdateScoresForAllPlayers(ctx, arg.GameID)
	if err != nil {
		return err
	}

	if winningCondition {
		s.EndGameHandler(ctx, arg.GameID, arg.PlayerGameID)
		log.Printf("Player %v has won the game", updatedPlayerGame.PlayerNumber)
	}

	return nil
}

// appendMissingIngredientWithWildcard replaces one missing ingredient using a wildcard.
func (s *GameService) appendMissingIngredientWithWildcard(ctx context.Context, ingredientsList []db.Card, requiredIngredients []string) ([]db.Card, error) {
	ingredientMap := make(map[string]bool)
	for _, ingredient := range ingredientsList {
		ingredientMap[ingredient.Name] = true
	}

	for _, required := range requiredIngredients {
		if !ingredientMap[required] {
			requiredCard, err := s.store.GetCardByName(ctx, required)
			if err != nil {
				return ingredientsList, fmt.Errorf("could not retrieve wildcard ingredient card from the database: %w", err)
			}
			ingredientsList = append(ingredientsList, requiredCard)
			break // Substitute for one missing ingredient only
		}
	}

	return ingredientsList, nil
}

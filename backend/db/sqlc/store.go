package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/PlatosCodes/desserted/backend/util"
)

// Store provides all functions to execure db queries and transactions
// Uses composition and extending the functionality of queries for single db operations
type Store interface {
	Querier
	RegisterTx(ctx context.Context, arg CreateUserParams) (RegisterTxResult, error)
	StartGameTx(ctx context.Context, arg StartGameTxParams) (StartGameTxResult, error)
	InitializeDeck(ctx context.Context, gameID int64, cardIDs []int64) (int64, error)
	DrawCard(ctx context.Context, arg DrawCardTxParams) (DrawCardTxResult, error)
	PlayDessertTx(ctx context.Context, arg PlayDessertTxParams) (PlayerGame, error)
	RefreshPlayerPantryTx(ctx context.Context, playerGameID int64, cardID int64) error
	StealRandomCardFromPlayerTx(ctx context.Context, playerGameID int64, cardID int64) (StealRandomCardFromPlayerTxResult, error)
	EndTurnTx(ctx context.Context, gameID int64, playerGameID int64) (EndTurnTxResult, error)
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	db *sql.DB
	*Queries
}

// NewStore creates a new Store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// In future will add more params to extend functionality of RegisterTx
// type RegisterTxParams struct {
// 	Username       string `json:"username"`
// 	Password string `json:"password"`
// 	Email          string `json:"email"`
// **Other fields here**
// }

type RegisterTxResult struct {
	User User `json:"user"`
}

// RegisterTx performs a new user registration.
// It creates a new user only, so there is no reason to actually use
// this besides getting practice for now, and adding new
// multi-operation database transaction features later
// **RegisterTxResult is also rather useless for now, but will be useful when
// we have actual transcations occuring.
func (store *SQLStore) RegisterTx(ctx context.Context, arg CreateUserParams) (RegisterTxResult, error) {
	var result RegisterTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.User, err = q.CreateUser(ctx, arg)

		if err != nil {
			return err
		}

		// future additonal operations here
		// ex. result.StoreProfilePhoto, err = q.StoreProfilePhoto(...)..

		return nil
	})

	return result, err

}

// StartGameTxParams holds parameters for the StartGameTx function
type StartGameTxParams struct {
	GameID    int64   `json:"game_id"`
	CreatedBy int64   `json:"created_by"`
	PlayerIDs []int64 `json:"player_ids"`
	CardIDs   []int64 `json:"card_ids"`
}

// StartGameTxResult holds the result for the StartGameTx function
type StartGameTxResult struct {
	Game Game `json:"game"`
}

// StartGameTx starts a game and initializes the deck in a transaction
func (store *SQLStore) StartGameTx(ctx context.Context, arg StartGameTxParams) (StartGameTxResult, error) {
	var result StartGameTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		numberOfPlayers, gameID := int32(len(arg.PlayerIDs)), arg.GameID

		// Update game status to active
		err = q.StartGame(ctx, StartGameParams{
			NumberOfPlayers: numberOfPlayers,
			GameID:          gameID,
		})
		if err != nil {
			return err
		}

		// Initialize the deck with shuffled cards
		_, err = store.InitializeDeck(ctx, arg.GameID, arg.CardIDs)
		if err != nil {
			return err
		}

		var turn int32

		for index, playerID := range arg.PlayerIDs {
			turn = int32(index) + 1
			err = store.UpdatePlayerNumber(ctx, UpdatePlayerNumberParams{
				PlayerNumber: sql.NullInt32{
					Int32: int32(turn),
					Valid: true,
				},
				PlayerGameID: playerID,
			})
			if err != nil {
				return err
			}

			for i := 0; i < 7; i++ { // Deal 7 cards to each player
				top_card, err := store.DrawTopCard(ctx, arg.GameID)
				if err != nil {
					return fmt.Errorf("failed to draw top card: %w", err)
				}
				_, err = q.AddCardToPlayerHand(ctx, AddCardToPlayerHandParams{
					PlayerGameID: playerID,
					CardID:       top_card,
				})
				if err != nil {
					return fmt.Errorf("failed to add card to player's hand: %w", err)
				}
				err = store.RemoveCardFromDeck(ctx, RemoveCardFromDeckParams{
					GameID: arg.GameID,
					CardID: top_card,
				})
				if err != nil {
					return fmt.Errorf("failed to remove card from game deck: %w", err)
				}
			}

			err = store.CreatePlayerTurnActions(ctx, playerID)
			if err != nil {
				return fmt.Errorf("failed to create entry into player turn actions table: %w", err)
			}
		}

		// Retrieve the updated game details
		result.Game, err = q.GetGameByID(ctx, arg.GameID)
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

// InitializeDeck initializes the deck for a game.
func (store *SQLStore) InitializeDeck(ctx context.Context, gameID int64, cardIDs []int64) (int64, error) {
	game_deck_id := int64(0)

	// Begin transaction
	tx, err := store.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return game_deck_id, fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Shuffle cardIDs
	util.Rand().Seed(time.Now().UnixNano())
	util.Rand().Shuffle(len(cardIDs), func(i, j int) {
		cardIDs[i], cardIDs[j] = cardIDs[j], cardIDs[i]
	})

	// Insert each card into game_deck table with order_index
	for index, cardID := range cardIDs {
		game_deck_id, err = store.InsertIntoGameDeck(ctx, InsertIntoGameDeckParams{
			GameID:     gameID,
			CardID:     cardID,
			OrderIndex: int32(index),
		})
		if err != nil {
			tx.Rollback()
			return game_deck_id, fmt.Errorf("failed to insert card into game_deck: %w", err)
		}
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return game_deck_id, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return game_deck_id, nil
}

// RefreshPlayerHand discards the player's hand and draws the same number of new cards.
func (store *SQLStore) RefreshPlayerPantryTx(ctx context.Context, playerGameID int64, cardID int64) error {

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Check if player has already played a special card this turn
		hasAlreadyPlayedSpecialCard, err := store.CheckSpecialCardPlayed(ctx, playerGameID)
		if err != nil {
			return fmt.Errorf("failed to get player's drawn card status for this turn: %w", err)
		}

		if hasAlreadyPlayedSpecialCard {
			return fmt.Errorf("player has already played a special card this turn: %w", err)
		}

		// Fetch the gameID based on playerGameID
		game, err := store.GetGameByPlayerGameID(ctx, playerGameID)
		if err != nil {
			return fmt.Errorf("failed to get game with playerGameID: %w", err)
		}
		gameID := game.GameID

		// Check if card is in player's hand
		inHand, err := q.IsCardInPlayerHand(ctx, IsCardInPlayerHandParams{
			PlayerGameID: playerGameID,
			CardID:       cardID,
		})
		if err != nil || !inHand {
			return fmt.Errorf("card not in player's hand: %w", err)
		}

		// Record card played (not removing from hand because it will be removed when all cards removed below)
		err = q.RecordPlayedCard(ctx, RecordPlayedCardParams{
			PlayerGameID: playerGameID,
			CardID:       cardID,
		})
		if err != nil {
			return fmt.Errorf("error recording card played: %w", err)
		}

		// Fetch current player hand
		currentHand, err := store.GetPlayerHand(ctx, playerGameID)
		if err != nil {
			return err
		}

		// Remove all cards in hand
		for _, card := range currentHand {
			err := store.RemoveCardFromPlayerHand(ctx, RemoveCardFromPlayerHandParams{
				PlayerGameID: playerGameID,
				CardID:       card.CardID,
			})
			if err != nil {
				return err
			}
		}

		// Draw new cards equal to the number of discarded cards
		for i := 0; i < len(currentHand); i++ {
			topCard, err := store.DrawTopCard(ctx, gameID)
			if err != nil {
				return fmt.Errorf("failed to draw top card: %w", err)
			}
			_, err = store.AddCardToPlayerHand(ctx, AddCardToPlayerHandParams{
				PlayerGameID: playerGameID,
				CardID:       topCard,
			})
			if err != nil {
				return fmt.Errorf("failed to add card to player's hand: %w", err)
			}
			err = store.RemoveCardFromDeck(ctx, RemoveCardFromDeckParams{
				GameID: gameID,
				CardID: topCard,
			})
			if err != nil {
				return fmt.Errorf("failed to remove card from game deck: %w", err)
			}
		}

		// Update special card played status
		err = q.UpdateSpecialCardPlayedStatus(ctx, playerGameID)
		if err != nil {
			log.Printf("Error updating special card played status: %v", err)
		}

		return nil

	})
	log.Println(err)
	return err
}

type StealRandomCardFromPlayerTxResult struct {
	TargetPlayerID int64 `json:"target_player_id"`
	StolenCardID   int64 `json:"stolen_card_id"`
}

func (store *SQLStore) StealRandomCardFromPlayerTx(ctx context.Context, playerGameID int64, cardID int64) (StealRandomCardFromPlayerTxResult, error) {
	var targetPlayerID, stolenCardID int64

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Check if player has already played a special card this turn
		hasAlreadyPlayedSpecialCard, err := store.CheckSpecialCardPlayed(ctx, playerGameID)
		if err != nil {
			return fmt.Errorf("failed to get player's drawn card status for this turn: %w", err)
		}

		if hasAlreadyPlayedSpecialCard {
			return fmt.Errorf("player has already played a special card this turn: %w", err)
		}

		// Fetch the gameID based on playerGameID
		game, err := store.GetGameByPlayerGameID(ctx, playerGameID)
		if err != nil {
			return fmt.Errorf("failed to get game with playerGameID: %w", err)
		}
		gameID := game.GameID

		// Check if steal card is in player's hand
		inHand, err := q.IsCardInPlayerHand(ctx, IsCardInPlayerHandParams{
			PlayerGameID: playerGameID,
			CardID:       cardID,
		})
		if err != nil || !inHand {
			return fmt.Errorf("card not in player's hand: %w", err)
		}

		// Record steal card played and remove from player's hand
		err = q.RecordPlayedCard(ctx, RecordPlayedCardParams{
			PlayerGameID: playerGameID,
			CardID:       cardID,
		})
		if err != nil {
			return fmt.Errorf("error recording card played: %w", err)
		}

		err = q.RemoveCardFromPlayerHand(ctx, RemoveCardFromPlayerHandParams{
			PlayerGameID: playerGameID,
			CardID:       cardID,
		})
		if err != nil {
			return fmt.Errorf("error removing card from hand: %w", err)
		}

		// Fetch all player IDs in the current game except the player who played the Steal Card
		players, err := q.ListGamePlayers(ctx, ListGamePlayersParams{
			GameID: gameID,
			Limit:  100,
			Offset: 0,
		})
		if err != nil {
			return fmt.Errorf("error fetching other player IDs: %w", err)
		}

		var playerIDs []int64

		for _, player := range players {
			if player.PlayerGameID != playerGameID {
				playerIDs = append(playerIDs, player.PlayerGameID)
			}
		}

		// Select a random player from the game
		if len(playerIDs) == 0 {
			return errors.New("no other players to steal from")
		}
		targetPlayerID = playerIDs[util.Rand().Intn(len(playerIDs))]

		// Fetch a random card from the selected player's hand
		targetHand, err := q.GetPlayerHand(ctx, targetPlayerID)
		if err != nil {
			return fmt.Errorf("error fetching target player hand: %w", err)
		}
		if len(targetHand) == 0 {
			return errors.New("target player has no cards to steal")
		}
		stolenCardID = targetHand[util.Rand().Intn(len(targetHand))].CardID

		// Remove the stolen card from the target player's hand
		err = q.RemoveCardFromPlayerHand(ctx, RemoveCardFromPlayerHandParams{
			PlayerGameID: targetPlayerID,
			CardID:       stolenCardID,
		})
		if err != nil {
			return fmt.Errorf("error removing card from target player's hand: %w", err)
		}

		// Add the stolen card to the current player's hand
		_, err = q.AddCardToPlayerHand(ctx, AddCardToPlayerHandParams{
			PlayerGameID: playerGameID,
			CardID:       stolenCardID,
		})
		if err != nil {
			return fmt.Errorf("error adding card to current player's hand: %w", err)
		}

		// Update special card played status
		err = q.UpdateSpecialCardPlayedStatus(ctx, playerGameID)
		if err != nil {
			log.Printf("Error updating special card played status: %v", err)
		}

		return nil
	})

	return StealRandomCardFromPlayerTxResult{
		TargetPlayerID: targetPlayerID,
		StolenCardID:   stolenCardID,
	}, err
}

type EndTurnTxResult struct {
	Game Game `json:"game"`
}

// EndTurnTx ends the current turn and updates the game to the next turn
func (store *SQLStore) EndTurnTx(ctx context.Context, gameID int64, playerGameID int64) (EndTurnTxResult, error) {
	var result EndTurnTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		// Retrieve the current game state
		game, err := q.GetGameByID(ctx, gameID)
		if err != nil {
			return fmt.Errorf("error fetching game: %w", err)
		}

		// Reset current player turn action status for next turn
		err = q.ResetTurnActions(ctx, playerGameID)
		if err != nil {
			return fmt.Errorf("error resetting player turn actions: %w", err)
		}

		// Calculate the next turn and player number
		nextTurn := game.CurrentTurn + 1
		nextPlayerNumber := ((nextTurn - 1) % game.NumberOfPlayers) + 1

		// Update the game state
		err = q.UpdateGameState(ctx, UpdateGameStateParams{
			GameID:      gameID,
			CurrentTurn: nextTurn,
			CurrentPlayerNumber: sql.NullInt32{
				Int32: int32(nextPlayerNumber),
				Valid: true,
			},
		})
		if err != nil {
			return fmt.Errorf("error updating game state: %w", err)
		}

		// Retrieve updated game details
		result.Game, err = q.GetGameByID(ctx, gameID)
		if err != nil {
			return fmt.Errorf("error retrieving updated game: %w", err)
		}

		return nil
	})

	return result, err
}

package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/PlatosCodes/desserted/backend/util"
	"github.com/PlatosCodes/desserted/backend/val"
)

// Store provides all functions to execure db queries and transactions
// Uses composition and extending the functionality of queries for single db operations
type Store interface {
	Querier
	RegisterTx(ctx context.Context, arg CreateUserParams) (RegisterTxResult, error)
	StartGameTx(ctx context.Context, arg StartGameTxParams) (StartGameTxResult, error)
	InitializeDeck(ctx context.Context, gameID int64, cardIDs []int64) (int64, error)
	DrawCard(ctx context.Context, arg DrawCardTxParams) (int64, error)
	PlayDessertTx(ctx context.Context, arg PlayDessertTxParams) (PlayDessertTxResult, error)
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
	PlayerIDs []int64 `json:"player_ids"`
	CardIDs   []int64 `json:"card_ids"` // Add this to include the card IDs for initializing the deck
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

		// Update game status to active
		err = q.UpdateGameStatus(ctx, UpdateGameStatusParams{
			Status: "active",
			GameID: arg.GameID,
		})
		if err != nil {
			return err
		}

		// Initialize the deck with shuffled cards
		_, err = store.InitializeDeck(ctx, arg.GameID, arg.CardIDs)
		if err != nil {
			return err
		}

		for _, playerID := range arg.PlayerIDs {
			for i := 0; i < 7; i++ { // Deal 7 cards to each player
				top_card, err := store.DrawTopCard(ctx, arg.GameID)
				if err != nil {
					return fmt.Errorf("failed to draw top card: %w", err)
				}
				err = q.AddCardToPlayerHand(ctx, AddCardToPlayerHandParams{
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

// DrawTxParams holds parameters for the StartGameTx function
type DrawCardTxParams struct {
	GameID   int64 `json:"game_id"`
	PlayerID int64 `json:"player_id"`
}

// DrawCard draws the top card from the deck for a given game.
func (store *SQLStore) DrawCard(ctx context.Context, arg DrawCardTxParams) (int64, error) {
	var cardID int64

	// Begin transaction
	tx, err := store.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return 0, fmt.Errorf("failed to begin transaction: %w", err)
	}

	log.Println("Game ID: ", arg.GameID)
	// Draw the top card from the game deck
	cardID, err = store.DrawTopCard(ctx, arg.GameID)
	if err != nil {
		return 0, fmt.Errorf("failed to draw top card: %w", err)
	}

	err = store.AddCardToPlayerHand(ctx, AddCardToPlayerHandParams{
		PlayerGameID: arg.PlayerID,
		CardID:       cardID,
	})
	if err != nil {
		return 0, fmt.Errorf("failed to add card to player's hand: %w", err)
	}

	// Remove the drawn card from the game deck
	err = store.RemoveCardFromDeck(ctx, RemoveCardFromDeckParams{
		GameID: arg.GameID,
		CardID: cardID,
	})
	if err != nil {
		return 0, fmt.Errorf("failed to remove card from game deck: %w", err)
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return cardID, nil
}

type PlayDessertTxParams struct {
	PlayerGameID int64
	DessertName  string
	CardIDs      []int64
}

type PlayDessertTxResult struct {
	DessertPlayedID int64
	PlayerGame      PlayerGame
	GameOver        bool
}

func (store *SQLStore) PlayDessertTx(ctx context.Context, arg PlayDessertTxParams) (PlayDessertTxResult, error) {
	var result PlayDessertTxResult

	log.Println("Dessert args:", arg)

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		ingredientsList := []string{}
		// Validate and process each card
		for _, cardID := range arg.CardIDs {
			// Fetch each card and validate
			card, err := q.GetCardByID(ctx, cardID)
			if err != nil {
				return fmt.Errorf("error fetching card: %w", err)
			}

			// Check if card is in player's hand
			inHand, err := q.IsCardInPlayerHand(ctx, IsCardInPlayerHandParams{
				PlayerGameID: arg.PlayerGameID,
				CardID:       cardID,
			})
			if err != nil || !inHand {
				return fmt.Errorf("card not in player's hand: %w", err)
			}

			// Record card played and remove from player's hand
			err = q.RecordPlayedCard(ctx, RecordPlayedCardParams{
				PlayerGameID: arg.PlayerGameID,
				CardID:       cardID,
			})
			if err != nil {
				return fmt.Errorf("error recording card played: %w", err)
			}

			err = q.RemoveCardFromPlayerHand(ctx, RemoveCardFromPlayerHandParams{
				PlayerGameID: arg.PlayerGameID,
				CardID:       cardID,
			})
			if err != nil {
				return fmt.Errorf("error removing card from hand: %w", err)
			}

			ingredientsList = append(ingredientsList, card.Name)
		}

		// Validate the dessert
		err = val.ValidateDessert(arg.DessertName, ingredientsList)
		if err != nil {
			log.Printf("Invalid dessert: %v", err)
			// Send error back to client
			return fmt.Errorf("invalid dessert: %v", err)
		}

		// Record the dessert played and update player's score
		dessert, err := q.GetDessertByName(ctx, arg.DessertName)
		if err != nil {
			return fmt.Errorf("error fetching dessert: %w", err)
		}

		err = q.RecordDessertPlayed(ctx, RecordDessertPlayedParams{
			PlayerGameID: arg.PlayerGameID,
			DessertID:    dessert.DessertID,
		})
		if err != nil {
			return fmt.Errorf("error recording dessert played: %w", err)
		}

		currPlayer, err := q.GetPlayerGame(ctx, arg.PlayerGameID)
		if err != nil {
			return fmt.Errorf("error getting player's previous score: %w", err)
		}

		updatedPlayerGame, err := q.UpdatePlayerScore(ctx, UpdatePlayerScoreParams{
			PlayerGameID: arg.PlayerGameID,
			PlayerScore:  sql.NullInt32{Int32: currPlayer.PlayerScore.Int32 + dessert.Points, Valid: true},
		})
		if err != nil {
			return fmt.Errorf("error updating player's score: %w", err)
		}

		// Check if the game is won
		winningCondition, err := q.IsGameWon(ctx, arg.PlayerGameID)
		if err != nil {
			return fmt.Errorf("error checking if game was won: %w", err)
		}

		// Set results
		result.DessertPlayedID = dessert.DessertID
		result.PlayerGame = updatedPlayerGame
		result.GameOver = winningCondition.Bool

		return nil
	})

	return result, err
}

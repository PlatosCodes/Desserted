package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/PlatosCodes/desserted/backend/util"
)

// Store provides all functions to execure db queries and transactions
// Uses composition and extending the functionality of queries for single db operations
type Store interface {
	Querier
	RegisterTx(ctx context.Context, arg CreateUserParams) (RegisterTxResult, error)
	StartGameTx(ctx context.Context, arg StartGameTxParams) (StartGameTxResult, error) // Add this line
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

type StartGameTxParams struct {
	GameID    int64   `json:"game_id"`
	PlayerIDs []int64 `json:"player_ids"`
}

type StartGameTxResult struct {
	Game Game `json:"game"`
}

// StartGameTx starts a game and deals cards to each player in a transaction
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

		// Shuffle and deal cards to each player
		cards, err := q.ListCards(ctx)
		if err != nil {
			return err
		}

		util.Rand().Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })

		for _, playerID := range arg.PlayerIDs {
			for i := 0; i < 7; i++ { // Deal 7 cards to each player
				err = q.AddCardToPlayerHand(ctx, AddCardToPlayerHandParams{
					PlayerGameID: playerID,
					CardID:       cards[i].CardID,
				})
				if err != nil {
					return err
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

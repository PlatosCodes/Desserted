// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: player_turns.sql

package db

import (
	"context"
	"database/sql"
)

const checkAllActionsCompleted = `-- name: CheckAllActionsCompleted :one
SELECT (card_drawn AND dessert_played AND special_card_played) AS all_actions_completed
FROM player_turn_actions
WHERE player_game_id = $1
`

// Checks if all actions for a turn are completed for a player
func (q *Queries) CheckAllActionsCompleted(ctx context.Context, playerGameID int64) (sql.NullBool, error) {
	row := q.db.QueryRowContext(ctx, checkAllActionsCompleted, playerGameID)
	var all_actions_completed sql.NullBool
	err := row.Scan(&all_actions_completed)
	return all_actions_completed, err
}

const checkCardDrawn = `-- name: CheckCardDrawn :one
SELECT (card_drawn)
FROM player_turn_actions
WHERE player_game_id = $1
`

// Checks if draw card action for a turn has been completed for a player
func (q *Queries) CheckCardDrawn(ctx context.Context, playerGameID int64) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkCardDrawn, playerGameID)
	var card_drawn bool
	err := row.Scan(&card_drawn)
	return card_drawn, err
}

const checkDessertPlayed = `-- name: CheckDessertPlayed :one
SELECT (dessert_played)
FROM player_turn_actions
WHERE player_game_id = $1
`

// Checks if play dessert card action for a turn has been completed for a player
func (q *Queries) CheckDessertPlayed(ctx context.Context, playerGameID int64) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkDessertPlayed, playerGameID)
	var dessert_played bool
	err := row.Scan(&dessert_played)
	return dessert_played, err
}

const checkSpecialCardPlayed = `-- name: CheckSpecialCardPlayed :one
SELECT (special_card_played)
FROM player_turn_actions
WHERE player_game_id = $1
`

// Checks if play special card action for a turn has been completed for a player
func (q *Queries) CheckSpecialCardPlayed(ctx context.Context, playerGameID int64) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkSpecialCardPlayed, playerGameID)
	var special_card_played bool
	err := row.Scan(&special_card_played)
	return special_card_played, err
}

const createPlayerTurnActions = `-- name: CreatePlayerTurnActions :exec
INSERT INTO player_turn_actions (player_game_id) 
VALUES ($1)
`

func (q *Queries) CreatePlayerTurnActions(ctx context.Context, playerGameID int64) error {
	_, err := q.db.ExecContext(ctx, createPlayerTurnActions, playerGameID)
	return err
}

const resetTurnActions = `-- name: ResetTurnActions :exec
UPDATE player_turn_actions
SET card_drawn = FALSE, dessert_played = FALSE, special_card_played = FALSE
WHERE player_game_id = $1
`

// Resets the turn actions for a player after their turn
func (q *Queries) ResetTurnActions(ctx context.Context, playerGameID int64) error {
	_, err := q.db.ExecContext(ctx, resetTurnActions, playerGameID)
	return err
}

const updateCardDrawnStatus = `-- name: UpdateCardDrawnStatus :exec
UPDATE player_turn_actions
SET card_drawn = TRUE
WHERE player_game_id = $1
`

// Updates the card drawn status for a player
func (q *Queries) UpdateCardDrawnStatus(ctx context.Context, playerGameID int64) error {
	_, err := q.db.ExecContext(ctx, updateCardDrawnStatus, playerGameID)
	return err
}

const updateDessertPlayedStatus = `-- name: UpdateDessertPlayedStatus :exec
UPDATE player_turn_actions
SET dessert_played = TRUE
WHERE player_game_id = $1
`

// Updates the dessert played status for a player
func (q *Queries) UpdateDessertPlayedStatus(ctx context.Context, playerGameID int64) error {
	_, err := q.db.ExecContext(ctx, updateDessertPlayedStatus, playerGameID)
	return err
}

const updateSpecialCardPlayedStatus = `-- name: UpdateSpecialCardPlayedStatus :exec
UPDATE player_turn_actions
SET special_card_played = TRUE
WHERE player_game_id = $1
`

// Updates the special card played status for a player
func (q *Queries) UpdateSpecialCardPlayedStatus(ctx context.Context, playerGameID int64) error {
	_, err := q.db.ExecContext(ctx, updateSpecialCardPlayedStatus, playerGameID)
	return err
}

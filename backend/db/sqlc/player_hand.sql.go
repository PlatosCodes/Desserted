// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: player_hand.sql

package db

import (
	"context"
)

const addCardToPlayerHand = `-- name: AddCardToPlayerHand :exec
INSERT INTO player_hand (player_game_id, card_id) VALUES ($1, $2)
`

type AddCardToPlayerHandParams struct {
	PlayerGameID int32 `json:"player_game_id"`
	CardID       int64 `json:"card_id"`
}

func (q *Queries) AddCardToPlayerHand(ctx context.Context, arg AddCardToPlayerHandParams) error {
	_, err := q.db.ExecContext(ctx, addCardToPlayerHand, arg.PlayerGameID, arg.CardID)
	return err
}

const getPlayerHand = `-- name: GetPlayerHand :many
SELECT card_id FROM player_hand WHERE player_game_id = $1
`

func (q *Queries) GetPlayerHand(ctx context.Context, playerGameID int32) ([]int64, error) {
	rows, err := q.db.QueryContext(ctx, getPlayerHand, playerGameID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var card_id int64
		if err := rows.Scan(&card_id); err != nil {
			return nil, err
		}
		items = append(items, card_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const recordPlayedCard = `-- name: RecordPlayedCard :exec
INSERT INTO played_cards (player_game_id, card_id, play_time) VALUES ($1, $2, NOW())
`

type RecordPlayedCardParams struct {
	PlayerGameID int32 `json:"player_game_id"`
	CardID       int64 `json:"card_id"`
}

func (q *Queries) RecordPlayedCard(ctx context.Context, arg RecordPlayedCardParams) error {
	_, err := q.db.ExecContext(ctx, recordPlayedCard, arg.PlayerGameID, arg.CardID)
	return err
}

const removeCardFromPlayerHand = `-- name: RemoveCardFromPlayerHand :exec
DELETE FROM player_hand WHERE player_hand_id = $1
`

func (q *Queries) RemoveCardFromPlayerHand(ctx context.Context, playerHandID int32) error {
	_, err := q.db.ExecContext(ctx, removeCardFromPlayerHand, playerHandID)
	return err
}
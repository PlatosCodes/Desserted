// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: desserts.sql

package db

import (
	"context"
)

const getDessertsPlayedByPlayer = `-- name: GetDessertsPlayedByPlayer :many
SELECT dessert_id
FROM dessert_played 
WHERE player_game_id = $1
`

func (q *Queries) GetDessertsPlayedByPlayer(ctx context.Context, playerGameID int64) ([]int32, error) {
	rows, err := q.db.QueryContext(ctx, getDessertsPlayedByPlayer, playerGameID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int32{}
	for rows.Next() {
		var dessert_id int32
		if err := rows.Scan(&dessert_id); err != nil {
			return nil, err
		}
		items = append(items, dessert_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const recordDessertPlayed = `-- name: RecordDessertPlayed :exec
INSERT INTO dessert_played (player_game_id, dessert_id) 
VALUES ($1, $2)
`

type RecordDessertPlayedParams struct {
	PlayerGameID int64 `json:"player_game_id"`
	DessertID    int32 `json:"dessert_id"`
}

func (q *Queries) RecordDessertPlayed(ctx context.Context, arg RecordDessertPlayedParams) error {
	_, err := q.db.ExecContext(ctx, recordDessertPlayed, arg.PlayerGameID, arg.DessertID)
	return err
}

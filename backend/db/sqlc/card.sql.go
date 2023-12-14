// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: card.sql

package db

import (
	"context"
	"database/sql"
)

const getAllCards = `-- name: GetAllCards :many

SELECT id, type, name, points FROM cards
`

// card.sql
// Get all cards
func (q *Queries) GetAllCards(ctx context.Context) ([]Card, error) {
	rows, err := q.db.QueryContext(ctx, getAllCards)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Card{}
	for rows.Next() {
		var i Card
		if err := rows.Scan(
			&i.ID,
			&i.Type,
			&i.Name,
			&i.Points,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCardByID = `-- name: GetCardByID :one
SELECT id, type, name, points FROM cards 
WHERE id = $1
`

// Get card by ID
func (q *Queries) GetCardByID(ctx context.Context, id int64) (Card, error) {
	row := q.db.QueryRowContext(ctx, getCardByID, id)
	var i Card
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Name,
		&i.Points,
	)
	return i, err
}

const getCardsByType = `-- name: GetCardsByType :many
SELECT id, type, name, points FROM cards 
WHERE type = $1
`

// Get cards by type
func (q *Queries) GetCardsByType(ctx context.Context, type_ sql.NullString) ([]Card, error) {
	rows, err := q.db.QueryContext(ctx, getCardsByType, type_)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Card{}
	for rows.Next() {
		var i Card
		if err := rows.Scan(
			&i.ID,
			&i.Type,
			&i.Name,
			&i.Points,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

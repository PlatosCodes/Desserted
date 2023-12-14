// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: card.sql

package db

import (
	"context"
	"database/sql"
)

const getCardByID = `-- name: GetCardByID :one
SELECT id, type, name FROM cards 
WHERE id = $1
`

// Get card by ID
func (q *Queries) GetCardByID(ctx context.Context, id int64) (Card, error) {
	row := q.db.QueryRowContext(ctx, getCardByID, id)
	var i Card
	err := row.Scan(&i.ID, &i.Type, &i.Name)
	return i, err
}

const listCards = `-- name: ListCards :many

SELECT id, type, name FROM cards
`

// card.sql
// List all cards
func (q *Queries) ListCards(ctx context.Context) ([]Card, error) {
	rows, err := q.db.QueryContext(ctx, listCards)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Card{}
	for rows.Next() {
		var i Card
		if err := rows.Scan(&i.ID, &i.Type, &i.Name); err != nil {
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

const listCardsByType = `-- name: ListCardsByType :many
SELECT id, type, name FROM cards 
WHERE type = $1
`

// List cards by type
func (q *Queries) ListCardsByType(ctx context.Context, type_ sql.NullString) ([]Card, error) {
	rows, err := q.db.QueryContext(ctx, listCardsByType, type_)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Card{}
	for rows.Next() {
		var i Card
		if err := rows.Scan(&i.ID, &i.Type, &i.Name); err != nil {
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

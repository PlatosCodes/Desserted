// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: card.sql

package db

import (
	"context"
)

const getCardByID = `-- name: GetCardByID :one
SELECT card_id, type, name FROM cards 
WHERE card_id = $1
`

// Get card by ID
func (q *Queries) GetCardByID(ctx context.Context, cardID int64) (Card, error) {
	row := q.db.QueryRowContext(ctx, getCardByID, cardID)
	var i Card
	err := row.Scan(&i.CardID, &i.Type, &i.Name)
	return i, err
}

const getCardByName = `-- name: GetCardByName :one
SELECT card_id, type, name FROM cards 
WHERE name = $1
`

// Get card by Name
func (q *Queries) GetCardByName(ctx context.Context, name string) (Card, error) {
	row := q.db.QueryRowContext(ctx, getCardByName, name)
	var i Card
	err := row.Scan(&i.CardID, &i.Type, &i.Name)
	return i, err
}

const listCardIDs = `-- name: ListCardIDs :many
SELECT card_id FROM cards
`

// List all cards
func (q *Queries) ListCardIDs(ctx context.Context) ([]int64, error) {
	rows, err := q.db.QueryContext(ctx, listCardIDs)
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

const listCards = `-- name: ListCards :many
SELECT card_id, type, name FROM cards
`

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
		if err := rows.Scan(&i.CardID, &i.Type, &i.Name); err != nil {
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
SELECT card_id, type, name FROM cards 
WHERE type = $1
`

// List cards by type
func (q *Queries) ListCardsByType(ctx context.Context, type_ string) ([]Card, error) {
	rows, err := q.db.QueryContext(ctx, listCardsByType, type_)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Card{}
	for rows.Next() {
		var i Card
		if err := rows.Scan(&i.CardID, &i.Type, &i.Name); err != nil {
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

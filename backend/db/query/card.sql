-- List all cards
-- name: ListCards :many
SELECT * FROM cards;

-- List all cards
-- name: ListCardIDs :many
SELECT card_id FROM cards;

-- Get card by ID
-- name: GetCardByID :one
SELECT * FROM cards 
WHERE card_id = $1;

-- List cards by type
-- name: ListCardsByType :many
SELECT * FROM cards 
WHERE type = $1;

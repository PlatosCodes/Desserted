-- card.sql

-- List all cards
-- name: ListCards :many
SELECT * FROM cards;

-- Get card by ID
-- name: GetCardByID :one
SELECT * FROM cards 
WHERE id = $1;

-- List cards by type
-- name: ListCardsByType :many
SELECT * FROM cards 
WHERE type = $1;

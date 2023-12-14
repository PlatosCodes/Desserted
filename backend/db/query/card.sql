-- card.sql

-- Get all cards
-- name: GetAllCards :many
SELECT * FROM cards;

-- Get card by ID
-- name: GetCardByID :one
SELECT * FROM cards 
WHERE id = $1;

-- Get cards by type
-- name: GetCardsByType :many
SELECT * FROM cards 
WHERE type = $1;

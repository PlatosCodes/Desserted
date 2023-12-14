-- Create a new game session
-- name: CreateGame :one
INSERT INTO games (created_by)
VALUES ($1)
RETURNING *;

-- End game
-- name: EndGame :exec
UPDATE games 
SET status = 'complete', ended_at = NOW() 
WHERE id = $1;

-- Get game session by ID
-- name: GetGameByID :one
SELECT * FROM games WHERE id = $1;

-- Draw a card
-- name: DrawCard :one
SELECT * FROM cards
ORDER BY RANDOM()
LIMIT 1;

-- Play a dessert
-- name: PlayDessert :exec
UPDATE players 
SET played_cards = array_append(played_cards, $1) 
WHERE user_id = $2 AND game_id = $3;

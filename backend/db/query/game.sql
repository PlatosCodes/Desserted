-- name: CreateGame :one
INSERT INTO games (created_by) 
VALUES ($1) 
RETURNING *;

-- name: GetGameByID :one
SELECT * FROM games 
WHERE game_id = $1;

-- name: ListActiveGames :many
SELECT * FROM games 
WHERE status = 'active' 
LIMIT $1 OFFSET $2;

-- name: UpdateGameStatus :exec
UPDATE games SET status = $1 
WHERE game_id = $2;

-- name: DeclareWinner :one
-- Declare the winner of the game
SELECT player_id FROM player_game WHERE game_id = ? ORDER BY player_score DESC LIMIT 1;

-- name: EndGame :exec
UPDATE games SET end_time = NOW() 
WHERE game_id = $1;

-- name: CreateGame :one
INSERT INTO games (created_by) 
VALUES ($1) 
RETURNING *;

-- name: GetGameByID :one
SELECT * FROM games 
WHERE game_id = $1;

-- name: ListGamePlayers :many
SELECT * FROM player_game 
WHERE game_id = $1 
LIMIT $2 OFFSET $3;

-- name: ListActiveGames :many
SELECT * FROM games 
LIMIT $1 OFFSET $2;

-- name: StartGame :exec
UPDATE games SET status = "active" 
WHERE game_id = $1;

-- name: UpdateGameStatus :exec
UPDATE games SET status = $1 
WHERE game_id = $2;

-- name: DeclareWinner :one
-- Declare the winner of the game
SELECT player_id FROM player_game 
WHERE game_id = $1 
ORDER BY player_score DESC LIMIT 1;

-- name: EndGame :exec
UPDATE games SET status = 'completed', end_time = NOW() 
WHERE game_id = $1;

-- name: AddPlayerToGame :exec
INSERT INTO player_game (player_id, game_id) 
VALUES ($1, $2);

-- name: GetPlayerGame :one
SELECT * FROM player_game 
WHERE player_game_id = $1;

-- name: UpdatePlayerGame :exec
UPDATE player_game 
SET player_score = $1, player_status = $2 
WHERE player_game_id = $3;

-- Check if a player has reached the winning condition
-- name: CheckWinCondition :one
SELECT player_id, player_score FROM player_game 
WHERE player_game_id = $1 AND player_score >= $2;

-- name: ListPlayerGames :many
SELECT * FROM player_game 
WHERE player_id = $1;

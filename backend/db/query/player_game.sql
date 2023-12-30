-- name: AddPlayerToGame :exec
INSERT INTO player_game (player_id, game_id) 
VALUES ($1, $2);

-- name: GetPlayerGame :one
SELECT * FROM player_game 
WHERE player_game_id = $1;

-- name: UpdatePlayerNumber :exec
UPDATE player_game 
SET player_number = $1
WHERE player_game_id = $2
;

-- name: UpdatePlayerScore :one
UPDATE player_game 
SET player_score = $1
WHERE player_game_id = $2
RETURNING *;

-- name: UpdatePlayerStatus :exec
UPDATE player_game 
SET player_status = $1
WHERE player_game_id = $2
RETURNING *;
;

-- Check if a player has reached the winning condition
-- name: IsGameWon :one
SELECT EXISTS (
  SELECT 1 FROM player_game
  WHERE player_score >= 100 AND player_game_id = $1
) OR NOT EXISTS (
  SELECT 1 FROM game_deck
  WHERE game_id = $1
) AS is_game_won;

-- name: ListPlayerGames :many
SELECT * FROM player_game 
WHERE player_id = $1;

-- name: ListActivePlayerGames :many
SELECT 
    player_game.player_game_id, 
    player_game.player_id, 
    player_game.game_id,
    games.number_of_players,
    player_game.player_number,
    player_game.player_score, 
    player_game.player_status,
    games.status, 
    games.created_by,
    games.current_turn, 
    games.current_player_number
FROM player_game 
INNER JOIN games ON player_game.game_id = games.game_id
WHERE player_id = $1 AND (games.status = 'active' OR games.status = 'waiting');

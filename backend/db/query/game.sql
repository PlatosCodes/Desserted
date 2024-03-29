-- name: CreateGame :one
INSERT INTO games (created_by) 
VALUES ($1) 
RETURNING *;

-- name: GetGameByID :one
SELECT * FROM games 
WHERE game_id = $1;

-- name: GetGameByPlayerGameID :one
SELECT * FROM games 
INNER JOIN 
  player_game ON games.game_id = player_game.game_id
WHERE 
  player_game.player_game_id = $1;

-- name: ListGamePlayers :many
SELECT * FROM player_game 
WHERE game_id = $1 
ORDER BY player_number ASC
LIMIT $2 OFFSET $3;

-- name: ListActiveGames :many
SELECT * FROM games 
LIMIT $1 OFFSET $2;

-- name: StartGame :exec
UPDATE games
SET status = 'active', number_of_players = $1, current_turn = 1, current_player_number = 1
WHERE game_id = $2;

-- name: UpdateGameState :exec
UPDATE games
SET current_turn = $2, current_player_number = $3, last_action_time = NOW()
WHERE game_id = $1;

-- name: UpdateGameStatus :exec
UPDATE games SET status = $1 
WHERE game_id = $2;

-- name: DeclareWinner :one
-- Declare the winner of the game
SELECT player_game_id, player_number, player_score FROM player_game 
WHERE game_id = $1 
ORDER BY player_score DESC LIMIT 1;

-- name: EndGame :exec
UPDATE games SET status = 'completed', end_time = NOW() 
WHERE game_id = $1;

-- name: GetGameScores :many
SELECT 
  users.id AS id,
  users.username,
  player_game.player_score
FROM 
  player_game
INNER JOIN 
  users ON player_game.player_id = users.id
WHERE 
  player_game.game_id = $1;

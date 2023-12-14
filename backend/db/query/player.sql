-- player.sql

-- Add player to a game
-- name: AddPlayerToGame :exec
INSERT INTO players (user_id, game_id, score, hand_cards, played_cards)
VALUES ($1, $2, $3, $4, $5);

-- Update player's score
-- name: UpdatePlayerScore :exec
UPDATE players SET score = $1 WHERE user_id = $2 AND game_id = $3;

-- Get players in a game
-- name: GetPlayersInGame :many
SELECT * FROM players WHERE game_id = $1;

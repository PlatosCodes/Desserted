-- name: InsertIntoGameDeck :one
INSERT INTO game_deck (game_id, card_id, order_index) 
VALUES ($1, $2, $3)
RETURNING game_deck_id;

-- name: GetGameDeck :one
SELECT * from game_deck
WHERE game_id = $1;

-- name: DrawTopCard :one
SELECT card_id FROM game_deck
WHERE game_id = $1
ORDER BY order_index ASC
LIMIT 1;

-- name: RemoveCardFromDeck :exec
DELETE FROM game_deck
WHERE game_id = $1 and card_id = $2;

-- name: IsDeckEmpty :one
SELECT NOT EXISTS (
  SELECT 1 FROM game_deck
  WHERE game_id = $1
) AS is_empty;


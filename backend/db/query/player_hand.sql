-- name: AddCardToPlayerHand :exec
INSERT INTO player_hand (player_game_id, card_id) 
VALUES ($1, $2);

-- name: RemoveCardFromPlayerHand :exec
DELETE FROM player_hand 
WHERE player_game_id = $1 and card_id = $2;

-- name: GetPlayerHand :many
SELECT card_id 
FROM player_hand 
WHERE player_game_id = $1;

-- name: RecordPlayedCard :exec
INSERT INTO played_cards (player_game_id, card_id) 
VALUES ($1, $2);

-- name: GetPlayedCards :many
SELECT * FROM played_cards
WHERE player_game_id = $1;
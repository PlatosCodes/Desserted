-- name: AddCardToPlayerHand :exec
INSERT INTO player_hand (player_game_id, card_id) 
VALUES ($1, $2);

-- name: RemoveCardFromPlayerHand :exec
DELETE FROM player_hand 
WHERE player_game_id = $1 and card_id = $2;

-- name: GetPlayerHand :many
SELECT player_hand.player_hand_id, player_hand.player_game_id, player_hand.card_id, cards.name
FROM player_hand 
JOIN cards ON player_hand.card_id = cards.card_id
WHERE player_hand.player_game_id = $1;

-- name: RecordPlayedCard :exec
INSERT INTO played_cards (player_game_id, card_id) 
VALUES ($1, $2);

-- name: GetPlayedCards :many
SELECT * FROM played_cards
WHERE player_game_id = $1;

-- name: IsCardInPlayerHand :one
SELECT EXISTS (
  SELECT 1 FROM player_hand
  WHERE player_game_id = $1 and card_id = $2
) AS in_hand;
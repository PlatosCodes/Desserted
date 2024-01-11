-- name: CreatePlayerTurnActions :exec
INSERT INTO player_turn_actions (player_game_id) 
VALUES ($1);

-- Updates the card drawn status for a player
-- name: UpdateCardDrawnStatus :exec
UPDATE player_turn_actions
SET card_drawn = TRUE
WHERE player_game_id = $1;

-- Updates the dessert played status for a player
-- name: UpdateDessertPlayedStatus :exec
UPDATE player_turn_actions
SET dessert_played = TRUE
WHERE player_game_id = $1;

-- Updates the special card played status for a player
-- name: UpdateSpecialCardPlayedStatus :exec
UPDATE player_turn_actions
SET special_card_played = TRUE
WHERE player_game_id = $1;

-- Checks if all actions for a turn are completed for a player
-- name: CheckAllActionsCompleted :one
SELECT (card_drawn AND dessert_played AND special_card_played) AS all_actions_completed
FROM player_turn_actions
WHERE player_game_id = $1;

-- Checks if draw card action for a turn has been completed for a player
-- name: CheckCardDrawn :one
SELECT (card_drawn)
FROM player_turn_actions
WHERE player_game_id = $1;

-- Checks if play dessert card action for a turn has been completed for a player
-- name: CheckDessertPlayed :one
SELECT (dessert_played)
FROM player_turn_actions
WHERE player_game_id = $1;

-- Checks if play special card action for a turn has been completed for a player
-- name: CheckSpecialCardPlayed :one
SELECT (special_card_played)
FROM player_turn_actions
WHERE player_game_id = $1;

-- Resets the turn actions for a player after their turn
-- name: ResetTurnActions :exec
UPDATE player_turn_actions
SET card_drawn = FALSE, dessert_played = FALSE, special_card_played = FALSE
WHERE player_game_id = $1;

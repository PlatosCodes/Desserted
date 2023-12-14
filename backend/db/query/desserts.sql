-- name: RecordDessertPlayed :exec
INSERT INTO dessert_played (player_game_id, dessert_id, icon_path, timestamp) 
VALUES ($1, $2, $3, NOW());

-- name: GetDessertsPlayedByPlayer :many
SELECT dessert_id, icon_path FROM dessert_played WHERE player_game_id = $1;

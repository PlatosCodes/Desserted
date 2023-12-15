-- name: RecordDessertPlayed :exec
INSERT INTO dessert_played (player_game_id, dessert_id) 
VALUES ($1, $2);

-- name: GetDessertsPlayedByPlayer :many
SELECT dessert_id
FROM dessert_played 
WHERE player_game_id = $1;


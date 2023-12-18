-- name: RecordDessertPlayed :exec
INSERT INTO dessert_played (player_game_id, dessert_id) 
VALUES ($1, $2);

-- name: GetDessertsPlayedByPlayer :many
SELECT dessert_id
FROM dessert_played 
WHERE player_game_id = $1;

-- name: GetDessertIDByName :one
SELECT dessert_id 
FROM desserts
WHERE name = $1;

-- name: GetDessertByName :one
SELECT * 
FROM desserts
WHERE name = $1;
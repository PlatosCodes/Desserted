-- name: CreateFriendship :one
INSERT INTO friends (friender_id, friendee_id) 
SELECT $1, $2
WHERE EXISTS (
  SELECT 1 FROM users WHERE id = $1
) AND EXISTS (
  SELECT 1 FROM users WHERE id = $2
)
RETURNING *;

-- name: ListUserFriends :many
SELECT * FROM friends
WHERE friender_id = $1 OR friendee_id = $1
ORDER BY friended_at
LIMIT $2
OFFSET $3;


-- name: GetFriendshipByID :one
SELECT * FROM friends
WHERE friendship_id = $1 LIMIT 1;

-- name: DeleteFriendship :exec
DELETE from friends WHERE friendship_id = $1;

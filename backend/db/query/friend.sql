-- name: CreateFriendship :one
INSERT INTO friends (friender_id, friendee_id)
SELECT $1, users.id
FROM users
WHERE users.username = $2 AND NOT EXISTS (
    SELECT 1 FROM friends
    WHERE friender_id = $1 AND friendee_id = users.id
)
RETURNING *;

-- name: ListUserFriends :many
SELECT * FROM friends
WHERE (friender_id = $1 OR friendee_id = $1) AND status = 'accepted'
ORDER BY friended_at
LIMIT $2
OFFSET $3;

-- name: ListPendingFriendRequests :many
SELECT users.id, users.username, friends.friendship_id, friends.friended_at
FROM friends
JOIN users ON friends.friender_id = users.id
WHERE friends.friendee_id = $1 AND friends.status = 'pending';

-- name: AcceptFriendRequest :exec
UPDATE friends
SET status = 'accepted', accepted_at = NOW()
WHERE friendee_id = $1 and friendship_id = $2 AND status = 'pending';

-- name: GetFriendshipByID :one
SELECT * FROM friends
WHERE friendship_id = $1 LIMIT 1;

-- name: DeleteFriendship :exec
DELETE from friends WHERE friendship_id = $1;

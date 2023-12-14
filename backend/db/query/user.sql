-- Register a new user
-- name: CreateUser :one
INSERT INTO users (username, password, email) 
VALUES ($1, $2, $3) 
RETURNING *;

-- Retrieve a user
-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- Get a user by their ID
-- name: GetUserByID :one
SELECT * FROM users 
WHERE id = $1;

-- Retrieve a list of all users
-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteUser :exec
DELETE from users WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
SET
  email = COALESCE(sqlc.narg(email), email),
  password = COALESCE(sqlc.narg(password), password),
  password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at)
WHERE
  username = sqlc.arg(username)
RETURNING *;
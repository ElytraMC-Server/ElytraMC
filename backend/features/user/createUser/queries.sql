-- name: CreateUser :one
INSERT INTO users (
  id, username, email
) VALUES (
  $1, $2, $3
)
RETURNING *;
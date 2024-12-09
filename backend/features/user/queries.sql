-- name: CreateUser :one
INSERT INTO users (
  id, username, email
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteUser :execrows
DELETE FROM users
WHERE id = $1;


-- name: GetUsers :many
SELECT * FROM users
ORDER BY username;



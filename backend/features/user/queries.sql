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

-- name: EditUser :execrows
UPDATE users 
SET username = $1, email = $2
WHERE id = $3;

-- name: LoginUser :one
SELECT * FROM users
WHERE username = $1 AND email = $2;
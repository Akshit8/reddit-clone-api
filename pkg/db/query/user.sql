-- name: CreateUser :one
INSERT INTO users (
  username,
  email,
  password
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: UpdateUserPassword :exec
UPDATE users SET password = $2 WHERE id = $1;

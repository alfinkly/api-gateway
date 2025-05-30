-- name: SelectUserById :one
SELECT id, username, name, surname, email
FROM users
WHERE id = $1;

-- name: SelectUserByEmail :one
SELECT id, username, name, surname, email
FROM users
WHERE email = $1;

-- name: CreateUser :one
INSERT INTO users (username, name, surname, email, password)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, username, name, surname, email, password;

-- name: UpdateUser :one
UPDATE users
SET username = $1, name = $2, surname = $3, email = $4, password = $5
WHERE id = $6
RETURNING id, username, name, surname, email, password;

-- name: DeleteUserById :exec
DELETE
FROM users
WHERE id = $1;

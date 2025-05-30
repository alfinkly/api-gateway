-- name: CreateUser :one
INSERT INTO users (id,
                   name,
                   email,
                   password)
VALUES ($1,
        $2,
        $3,
        $4
       )
RETURNING
    id,
    name,
    email,
    password;


-- name: GetUserByID :one
SELECT id,
       name,
       email,
       password
FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT id,
       name,
       email,
       password
FROM users
WHERE email = $1;

-- name: UpdateUser :one
UPDATE users
SET name  = $2,
    email = $3
WHERE id = $1
RETURNING
    id,
    name,
    email,
    password;

-- name: DeleteUser :exec
DELETE
FROM users
WHERE id = $1;

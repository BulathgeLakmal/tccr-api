-- name: CreateUser :one
INSERT INTO "user"(
    first_name,
    last_name,
    email,
    password,
    role
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;


-- name: GetUser :one
SELECT * FROM "user"
WHERE user_id = $1 LIMIT 1;

-- name: ListUser :many
SELECT * FROM "user"
WHERE email = $1
ORDER BY user_id
LIMIT $2
OFFSET $3;

-- name: UpdateUser :one
UPDATE "user"
SET email = $1
WHERE user_id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM "user"
WHERE user_id = $1;


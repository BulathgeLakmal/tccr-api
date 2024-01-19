-- name: CreateUser :one
INSERT INTO "user"(
    first_name,
    last_name,
    email,
    hashed_password,
    role,
    username
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;


-- name: GetUser :one
SELECT * FROM "user"
WHERE email = $1 LIMIT 1;




-- name: CreateUserDetails :one
INSERT INTO "userDetails"(
    user_id,
    phone,
    address_line1,
    address_line2
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;


-- name: GetUserDetails :one
SELECT * FROM "userDetails"
WHERE user_details_id = $1 LIMIT 1;

-- name: ListUserDetails :many
SELECT * FROM "userDetails"
WHERE user_details_id = $1
ORDER BY user_details_id
LIMIT $2
OFFSET $3;

-- name: UpdateUserDetails :one
UPDATE "userDetails"
SET address_line1 = $1
WHERE user_details_id = $1
RETURNING *;

-- name: DeleteUserDetails :exec
DELETE FROM "userDetails"
WHERE user_details_id = $1;


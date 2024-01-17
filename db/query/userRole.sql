-- name: CreateUserRole :one
INSERT INTO "userRole" (
    role
) VALUES (
    $1
)
RETURNING *;

-- name: GetUserRole :one
SELECT * FROM "userRole"
WHERE role_id = $1 LIMIT 1;

-- name: ListUserRole :many
SELECT * FROM "userRole"
WHERE role = $1
ORDER BY role_id
LIMIT $2
OFFSET $3;

-- name: UpdateUserRole :one
UPDATE "userRole"
SET role = $1
WHERE role_id = $1
RETURNING *;

-- name: DeleteUserRole :exec
DELETE FROM "userRole"
WHERE role_id = $1;


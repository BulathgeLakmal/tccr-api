-- name: CreateCategory :one
INSERT INTO category (
category_name,
category_desc
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetCategory :one
SELECT * FROM category
WHERE category_id = $1 LIMIT 1;

-- name: GetCategoryForUpdate :one
SELECT * FROM category
WHERE category_id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListCategory :many
SELECT * FROM category
WHERE category_name = $1
ORDER BY category_id
LIMIT $2
OFFSET $3;

-- name: UpdateCategory :one
UPDATE category
SET category_name = $1
WHERE category_id = $1
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM category
WHERE category_id = $1;

-- name: CreateCourseModule :one
INSERT INTO "courseModule"(
  course_id,
  module_name
) VALUES (
    $1, $2
)
RETURNING *;


-- name: GetCourseModule :one
SELECT * FROM "courseModule"
WHERE module_id = $1 LIMIT 1;


-- name: ListModules :many
SELECT * FROM "courseModule"
WHERE module_id = $1
ORDER BY course_id
LIMIT $2
OFFSET $3;

-- name: DeleteCourseModule :exec
DELETE FROM "courseModule"
WHERE module_id = $1;

-- name: UpdateCourseModule :one
UPDATE "courseModule"
SET module_name = $1
WHERE module_id = $1
RETURNING *;

-- name: GetAllCourseModule :many
SELECT * FROM "courseModule";
-- name: CreateAssignment :one
INSERT INTO "assignment"(
    user_id,
    course_module_id,
    lecture_id
) VALUES (
    $1, $2, $3
)
RETURNING *;


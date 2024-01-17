-- name: CreateAssignmentFile :one
INSERT INTO "assignment_file"(
    assignment_id,
    assignment_link
) VALUES (
    $1, $2
)
RETURNING *;


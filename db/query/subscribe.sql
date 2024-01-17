-- name: CreateSubscribe :one
INSERT INTO "subscribe"(
    user_id,
    course_id
) VALUES (
    $1, $2
)
RETURNING *;


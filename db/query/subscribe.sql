-- name: CreateSubscribe :one
INSERT INTO "subscribe"(
    user_id,
    course_id
) VALUES (
    $1, $2
)
RETURNING *;

-- name: GetSubscribe :one
SELECT * FROM "subscribe"
WHERE user_id = $1 LIMIT 1;
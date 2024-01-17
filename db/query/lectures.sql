-- name: CreateLecture :one
INSERT INTO "lectures"(
    course_module_id,
    lecture_desc,
    lecture_number,
    "video_URL",
    status
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;



-- name: GetLecture :one
SELECT * FROM "lectures"
WHERE lecture_id = $1 LIMIT 1;

-- name: ListLectures :many
SELECT * FROM "lectures"
WHERE lecture_number = $1
ORDER BY lecture_id
LIMIT $2
OFFSET $3;

-- name: UpdateLecture :one
UPDATE "lectures"
SET status = $1
WHERE lecture_id = $1
RETURNING *;

-- name: DeleteLecture :exec
DELETE FROM "lectures"
WHERE lecture_id = $1;


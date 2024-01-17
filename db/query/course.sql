-- name: CreateCourse :one
INSERT INTO course (
    course_name,
    course_desc,
    category
) VALUES(
    $1 , $2, $3
)
RETURNING *;


-- name: GetCourse :one
SELECT * FROM course
WHERE course_id = $1 LIMIT 1;

-- name: GetCourseForUpdate :one
SELECT * FROM course
WHERE course_id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListCourse :many
SELECT * FROM course
WHERE course_name = $1
ORDER BY course_id
LIMIT $2
OFFSET $3;

-- name: UpdateCourse :one
UPDATE course
SET course_name = $1
WHERE course_id = $1
RETURNING *;

-- name: DeleteCourse :exec
DELETE FROM course
WHERE course_id = $1;

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: subscribe.sql

package db

import (
	"context"
)

const createSubscribe = `-- name: CreateSubscribe :one
INSERT INTO "subscribe"(
    user_id,
    course_id
) VALUES (
    $1, $2
)
RETURNING subscribe_id, user_id, course_id
`

type CreateSubscribeParams struct {
	UserID   int64 `json:"user_id"`
	CourseID int64 `json:"course_id"`
}

func (q *Queries) CreateSubscribe(ctx context.Context, arg CreateSubscribeParams) (Subscribe, error) {
	row := q.db.QueryRowContext(ctx, createSubscribe, arg.UserID, arg.CourseID)
	var i Subscribe
	err := row.Scan(&i.SubscribeID, &i.UserID, &i.CourseID)
	return i, err
}

const getSubscribe = `-- name: GetSubscribe :one
SELECT subscribe_id, user_id, course_id FROM "subscribe"
WHERE user_id = $1 LIMIT 1
`

func (q *Queries) GetSubscribe(ctx context.Context, userID int64) (Subscribe, error) {
	row := q.db.QueryRowContext(ctx, getSubscribe, userID)
	var i Subscribe
	err := row.Scan(&i.SubscribeID, &i.UserID, &i.CourseID)
	return i, err
}

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: assignment_file.sql

package db

import (
	"context"
	"database/sql"
)

const createAssignmentFile = `-- name: CreateAssignmentFile :one
INSERT INTO "assignment_file"(
    assignment_id,
    assignment_link
) VALUES (
    $1, $2
)
RETURNING assignment_file_id, assignment_id, assignment_link
`

type CreateAssignmentFileParams struct {
	AssignmentID   sql.NullInt64 `json:"assignment_id"`
	AssignmentLink string        `json:"assignment_link"`
}

func (q *Queries) CreateAssignmentFile(ctx context.Context, arg CreateAssignmentFileParams) (AssignmentFile, error) {
	row := q.db.QueryRowContext(ctx, createAssignmentFile, arg.AssignmentID, arg.AssignmentLink)
	var i AssignmentFile
	err := row.Scan(&i.AssignmentFileID, &i.AssignmentID, &i.AssignmentLink)
	return i, err
}
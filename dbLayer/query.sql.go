// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: query.sql

package dbLayer

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createComment = `-- name: CreateComment :exec
INSERT INTO Comments (
    comment, publishedTime, urlString
)   VALUES (
    $1, $2, $3
)
`

type CreateCommentParams struct {
	Comment       pgtype.Text
	Publishedtime pgtype.Timestamp
	Urlstring     pgtype.Text
}

func (q *Queries) CreateComment(ctx context.Context, arg CreateCommentParams) error {
	_, err := q.db.Exec(ctx, createComment, arg.Comment, arg.Publishedtime, arg.Urlstring)
	return err
}

const createReplyComment = `-- name: CreateReplyComment :exec
INSERT INTO Comments (
    comment, publishedTime, parentCommentID, urlString
) VALUES (
    $1, $2, $3, $4
)
`

type CreateReplyCommentParams struct {
	Comment         pgtype.Text
	Publishedtime   pgtype.Timestamp
	Parentcommentid pgtype.Int4
	Urlstring       pgtype.Text
}

func (q *Queries) CreateReplyComment(ctx context.Context, arg CreateReplyCommentParams) error {
	_, err := q.db.Exec(ctx, createReplyComment,
		arg.Comment,
		arg.Publishedtime,
		arg.Parentcommentid,
		arg.Urlstring,
	)
	return err
}

const retrieveComments = `-- name: RetrieveComments :many
SELECT commentid, comment, publishedtime, parentcommentid, urlstring FROM Comments
WHERE parentCommentID IS NULL
`

func (q *Queries) RetrieveComments(ctx context.Context) ([]Comment, error) {
	rows, err := q.db.Query(ctx, retrieveComments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.Commentid,
			&i.Comment,
			&i.Publishedtime,
			&i.Parentcommentid,
			&i.Urlstring,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const retrieveReplies = `-- name: RetrieveReplies :many
SELECT commentid, comment, publishedtime, parentcommentid, urlstring FROM Comments
WHERE parentCommentID = $1
`

func (q *Queries) RetrieveReplies(ctx context.Context, parentcommentid pgtype.Int4) ([]Comment, error) {
	rows, err := q.db.Query(ctx, retrieveReplies, parentcommentid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.Commentid,
			&i.Comment,
			&i.Publishedtime,
			&i.Parentcommentid,
			&i.Urlstring,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
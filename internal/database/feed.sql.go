// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: feed.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeed = `-- name: CreateFeed :one
INSERT INTO gator.feeds 
(id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, created_at, updated_at, name, url, user_id, last_fetched_at
`

type CreateFeedParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Url       string
	UserID    uuid.UUID
}

func (q *Queries) CreateFeed(ctx context.Context, arg CreateFeedParams) (GatorFeed, error) {
	row := q.db.QueryRowContext(ctx, createFeed,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.Url,
		arg.UserID,
	)
	var i GatorFeed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Url,
		&i.UserID,
		&i.LastFetchedAt,
	)
	return i, err
}

const deleteFeeds = `-- name: DeleteFeeds :exec
DELETE FROM gator.feeds
`

func (q *Queries) DeleteFeeds(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteFeeds)
	return err
}

const getAllFeedNames = `-- name: GetAllFeedNames :many
SELECT name FROM gator.feeds
`

func (q *Queries) GetAllFeedNames(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getAllFeedNames)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFeedByURL = `-- name: GetFeedByURL :one
SELECT id, created_at, updated_at, name, url, user_id, last_fetched_at FROM gator.feeds
WHERE url = $1
`

func (q *Queries) GetFeedByURL(ctx context.Context, url string) (GatorFeed, error) {
	row := q.db.QueryRowContext(ctx, getFeedByURL, url)
	var i GatorFeed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Url,
		&i.UserID,
		&i.LastFetchedAt,
	)
	return i, err
}

const getFeedNameUrlUser = `-- name: GetFeedNameUrlUser :many
SELECT name, url, user_id 
FROM gator.feeds
`

type GetFeedNameUrlUserRow struct {
	Name   string
	Url    string
	UserID uuid.UUID
}

func (q *Queries) GetFeedNameUrlUser(ctx context.Context) ([]GetFeedNameUrlUserRow, error) {
	rows, err := q.db.QueryContext(ctx, getFeedNameUrlUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFeedNameUrlUserRow
	for rows.Next() {
		var i GetFeedNameUrlUserRow
		if err := rows.Scan(&i.Name, &i.Url, &i.UserID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getNextFeedToFetch = `-- name: GetNextFeedToFetch :one
SELECT id, created_at, updated_at, name, url, user_id, last_fetched_at from gator.feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1
`

func (q *Queries) GetNextFeedToFetch(ctx context.Context) (GatorFeed, error) {
	row := q.db.QueryRowContext(ctx, getNextFeedToFetch)
	var i GatorFeed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Url,
		&i.UserID,
		&i.LastFetchedAt,
	)
	return i, err
}

const markFeedFetched = `-- name: MarkFeedFetched :one
UPDATE gator.feeds 
SET last_fetched_at = NOW(), 
updated_at = NOW()
WHERE id = $1
RETURNING id, created_at, updated_at, name, url, user_id, last_fetched_at
`

func (q *Queries) MarkFeedFetched(ctx context.Context, id uuid.UUID) (GatorFeed, error) {
	row := q.db.QueryRowContext(ctx, markFeedFetched, id)
	var i GatorFeed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Url,
		&i.UserID,
		&i.LastFetchedAt,
	)
	return i, err
}

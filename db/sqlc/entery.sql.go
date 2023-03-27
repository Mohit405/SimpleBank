// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: entery.sql

package db

import (
	"context"
)

const createEntery = `-- name: CreateEntery :one
INSERT INTO enteries(
    account_id,
    amount
)VALUES(
    $1, $2
)RETURNING id, account_id, amount, created_at
`

type CreateEnteryParams struct {
	AccountID int64 `json:"account_id"`
	Amount    int64 `json:"amount"`
}

func (q *Queries) CreateEntery(ctx context.Context, arg CreateEnteryParams) (Entery, error) {
	row := q.db.QueryRowContext(ctx, createEntery, arg.AccountID, arg.Amount)
	var i Entery
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getEntry = `-- name: GetEntry :one
SELECT id, account_id, amount, created_at FROM enteries
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetEntry(ctx context.Context, id int64) (Entery, error) {
	row := q.db.QueryRowContext(ctx, getEntry, id)
	var i Entery
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listEntries = `-- name: ListEntries :many
SELECT id, account_id, amount, created_at FROM enteries
WHERE account_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListEntriesParams struct {
	AccountID int64 `json:"account_id"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *Queries) ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entery, error) {
	rows, err := q.db.QueryContext(ctx, listEntries, arg.AccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Entery
	for rows.Next() {
		var i Entery
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
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

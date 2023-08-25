// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: accounts.sql

package database

import (
	"context"
	"time"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts (
  id, name, balance, created_at, updated_at
  ) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING id, name, balance, created_at, updated_at
`

type CreateAccountParams struct {
	ID        int64
	Name      string
	Balance   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount,
		arg.ID,
		arg.Name,
		arg.Balance,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Balance,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAccount, id)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT id, name, balance, created_at, updated_at FROM accounts
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, id int64) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Balance,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listAccounts = `-- name: ListAccounts :many
SELECT id, name, balance, created_at, updated_at FROM accounts
ORDER BY name
`

func (q *Queries) ListAccounts(ctx context.Context) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, listAccounts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Account
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Balance,
			&i.CreatedAt,
			&i.UpdatedAt,
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
-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY name;

-- name: CreateAccount :one
INSERT INTO accounts (
  id, name, balance, created_at, updated_at
  ) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;
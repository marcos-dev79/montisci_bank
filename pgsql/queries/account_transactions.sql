-- name: GetAccountTransactions :one
SELECT * FROM account_transactions
WHERE id = $1 LIMIT 1;

-- name: ListAccountTransactions :many
SELECT * FROM account_transactions
ORDER BY id;

-- name: CreateAccountTransactions :one
INSERT INTO account_transactions (
  id, account_id, amount, created_at, updated_at
  ) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: DeleteAccountTransactions :exec
DELETE FROM account_transactions
WHERE id = $1;
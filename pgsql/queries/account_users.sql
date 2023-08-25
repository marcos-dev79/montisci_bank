-- name: GetAccountUsers :one
SELECT * FROM account_users
WHERE user_id = $1 LIMIT 1;

-- name: GetAccountUsersID :one
SELECT * FROM account_users
WHERE id = $1 LIMIT 1;

-- name: ListAccountUsers :many
SELECT * FROM account_users
ORDER BY created_at;

-- name: CreateAccountUser :one
INSERT INTO account_users (
  id, account_id, user_id, created_at, updated_at
  ) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: DeleteAccountUser :exec
DELETE FROM account_users
WHERE user_id = $1;
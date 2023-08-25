-- name: GetProfession :one
SELECT * FROM professions
WHERE id = $1 LIMIT 1;

-- name: ListProfessions :many
SELECT * FROM professions
ORDER BY name;

-- name: CreateProfession :one
INSERT INTO professions (
  id, name, created_at, updated_at
  ) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: DeleteProfession :exec
DELETE FROM professions
WHERE id = $1;
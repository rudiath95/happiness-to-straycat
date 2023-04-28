-- name: CreateTag :one
INSERT INTO tags (
  name
) VALUES (
  $1
) RETURNING *;

-- name: GetTagByID :one
SELECT * FROM tags
WHERE id = $1 LIMIT 1;

-- name: GetTagByName :one
SELECT * FROM tags
WHERE name = $1 LIMIT 1;

-- name: ListTags :many
SELECT * FROM tags
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteTag :exec
DELETE FROM tags 
WHERE id = $1;
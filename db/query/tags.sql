-- name: CreateTag :one
INSERT INTO tags (
  name,
  updated_at
  
) VALUES (
  $1,$2
) RETURNING *;

-- name: UpdateTag :one
UPDATE tags
set 
name = coalesce(sqlc.narg('name'), name), 
updated_at = coalesce(sqlc.narg('updated_at'), updated_at ) 
WHERE id = sqlc.arg('id')
RETURNING *;

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
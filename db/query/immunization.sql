-- name: CreateImmunization :one
INSERT INTO immunization (
  name,
  updated_at
  
) VALUES (
  $1,$2
) RETURNING *;

-- name: UpdateImmunization :one
UPDATE immunization
set 
name = coalesce(sqlc.narg('name'), name), 
updated_at = coalesce(sqlc.narg('updated_at'), updated_at ) 
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: GetImmunizationByID :one
SELECT * FROM immunization
WHERE id = $1 LIMIT 1;

-- name: GetImmunizationByName :one
SELECT * FROM immunization
WHERE name = $1 LIMIT 1;

-- name: ListImmunizations :many
SELECT * FROM immunization
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteImmunization :exec
DELETE FROM immunization 
WHERE id = $1;